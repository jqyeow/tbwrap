package tbwrap

//go:generate mockgen -destination=./mocks/mock_Bot.go -package=mocks github.com/enrico5b1b4/tbwrap Bot
//go:generate mockgen -destination=./mocks/mock_TeleBot.go -package=mocks github.com/enrico5b1b4/tbwrap TeleBot

import (
	"fmt"
	"log"
	"regexp"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type HandlerFunc func(c Context) error

type Route struct {
	Path    *regexp.Regexp
	Handler HandlerFunc
}

type TeleBot interface {
	Handle(endpoint interface{}, handler interface{})
	Respond(callback *tb.Callback, responseOptional ...*tb.CallbackResponse) error
	Send(to tb.Recipient, what interface{}, options ...interface{}) (*tb.Message, error)
	Start()
}

type Bot interface {
	Respond(callback *tb.Callback, responseOptional ...*tb.CallbackResponse) error
	Send(to tb.Recipient, what interface{}, options ...interface{}) (*tb.Message, error)
}

type TBWrapBot struct {
	TBot   TeleBot
	Routes map[*regexp.Regexp]*Route
}

type Config struct {
	Token         string
	AllowedUsers  []int
	AllowedGroups []int
}

func NewBot(cfg Config) (*TBWrapBot, error) {
	poller := NewPollerWithAllowedUserAndGroups(15*time.Second, cfg.AllowedUsers, cfg.AllowedGroups)
	tBot, err := tb.NewBot(tb.Settings{
		Token:  cfg.Token,
		Poller: poller,
	})
	if err != nil {
		return nil, err
	}

	b := &TBWrapBot{
		TBot:   tBot,
		Routes: map[*regexp.Regexp]*Route{},
	}

	// b.handle(tb.OnText, func(m *tb.Message) {
	// 	// all the text messages that weren't
	// 	// captured by existing handlers
	// 	b.HandleOnText(m.Text, m.Chat)
	// })

	return b, nil
}

func (b *TBWrapBot) handle(endpoint interface{}, handler interface{}) {
	b.TBot.Handle(endpoint, handler)
}

func (b *TBWrapBot) Respond(callback *tb.Callback, responseOptional ...*tb.CallbackResponse) error {
	return b.TBot.Respond(callback, responseOptional...)
}

func (b *TBWrapBot) Send(to tb.Recipient, what interface{}, options ...interface{}) (*tb.Message, error) {
	mergedOptions := append([]interface{}{&tb.SendOptions{ParseMode: tb.ModeMarkdown}}, options...)

	return b.TBot.Send(to, what, mergedOptions...)
}

func (b *TBWrapBot) AddRegExp(path string, handler HandlerFunc) {
	compiledRegExp := regexp.MustCompile(path)

	b.Routes[compiledRegExp] = &Route{Path: compiledRegExp, Handler: handler}
}

func (b *TBWrapBot) AddMultiRegExp(paths []string, handler HandlerFunc) {
	for i := range paths {
		compiledRegExp := regexp.MustCompile(paths[i])

		b.Routes[compiledRegExp] = &Route{Path: compiledRegExp, Handler: handler}
	}
}

func (b *TBWrapBot) Add(path string, handler HandlerFunc) {
	b.handle(path, func(m *tb.Message) {
		c := &context{chat: m.Chat, text: m.Text, chatID: int(m.Chat.ID), bot: b}
		err := handler(c)
		if err != nil {
			_ = c.Send(fmt.Sprintf("%s", err))
			log.Println(err)
		}
	})
}

func (b *TBWrapBot) AddButton(path *tb.InlineButton, handler HandlerFunc) {
	b.handle(path, func(callback *tb.Callback) {
		c := &context{
			chat:     callback.Message.Chat,
			text:     callback.Message.Text,
			callback: callback,
			chatID:   int(callback.Message.Chat.ID),
			bot:      b,
		}
		err := handler(c)
		if err != nil {
			_ = c.Send(fmt.Sprintf("%s", err))
			log.Println(err)
		}
	})
}

func (b *TBWrapBot) HandleOnText(text string, chat *tb.Chat) {
	for regExpKey := range b.Routes {
		matches := regExpKey.FindStringSubmatch(text)
		names := regExpKey.SubexpNames()

		if len(matches) > 0 {
			params := mapSubexpNames(matches, names)
			c := &context{chat: chat, text: text, params: params, chatID: int(chat.ID), route: regExpKey, bot: b}
			err := b.Routes[regExpKey].Handler(c)
			if err != nil {
				_ = c.Send(fmt.Sprintf("%s", err))
				log.Println(err)
			}
		}
	}
}

func (b *TBWrapBot) Start() {
	b.handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
		b.HandleOnText(m.Text, m.Chat)
	})

	b.TBot.Start()
}

func mapSubexpNames(m, n []string) map[string]string {
	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i := range n {
		r[n[i]] = m[i]
	}
	return r
}
