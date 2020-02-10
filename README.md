# tbwrap
Wrapper library for [github.com/tucnak/telebot](https://github.com/tucnak/telebot) to create simple text-based Telegram bots

## Installation

```
go get github.com/enrico5b1b4/tbwrap
```

## Examples

[Ping Bot](./examples/ping) - A bot which responds to your ping with a pong  
[Hello Bot](./examples/ping) - A bot which greets you  
[Todo List Bot](./examples/todo) - A bot that keeps a todo list for you  
[Joke Bot](./examples/joke) - A bot that sends you a joke

## Usage
### Simple
```go
package main

import (
	"log"

	"github.com/enrico5b1b4/tbwrap"
)

func main() {
	telegramBotToken := "TELEGRAM_BOT_TOKEN"

	botConfig := tbwrap.Config{
		Token: telegramBotToken,
	}
	telegramBot, err := tbwrap.NewBot(botConfig)
	if err != nil {
		log.Fatal(err)
	}

	telegramBot.Add(`/ping`, func(c tbwrap.Context) error {
		return c.Send("pong!")
	})
	telegramBot.Start()
}
```
### Private bot
You can make the bot respond only to certain chats (users or groups) by passing a list of ids.  

```go
// ...
botConfig := tbwrap.Config{
    Token:         telegramBotToken,
    AllowedUsers:  "123456,234567,345678", // only users with these ids can interact with the bot
    AllowedGroups: "-999999", // only groups with these ids can interact with the bot
}
telegramBot, err := tbwrap.NewBot(botConfig)
if err != nil {
    log.Fatal(err)
}
// ...
}
```
### Messages
Use regular expression named capturing groups to parametrise incoming messages
```go
// ...

type HelloMessage struct {
    Name string `regexpGroup:"name"`
}

func main() {
	// ...
	telegramBot.AddRegExp(`\/hello (?P<name>.*)`, func(c tbwrap.Context) error {
		message := new(HelloMessage)
		if err := c.Bind(message); err != nil {
			return err
		}

		return c.Send(fmt.Sprintf("Hello %s!", message.Name))
    })
	// ...
}
```