package main

import (
	"fmt"

	"github.com/enrico5b1b4/tbwrap"
)

type RemoveMessage struct {
	Index int `regexpGroup:"index"`
}

func HandleRemove(todos map[int][]string) func(c tbwrap.Context) error {
	return func(c tbwrap.Context) error {
		message := new(RemoveMessage)
		if err := c.Bind(message); err != nil {
			return err
		}

		chatTodos := todos[c.ChatID()]
		if message.Index < 0 || message.Index > len(chatTodos)-1 {
			return c.Send(fmt.Sprintf(`cannot remove entry "%d"`, message.Index))
		}

		value := todos[c.ChatID()][message.Index]
		todos[c.ChatID()] = append(todos[c.ChatID()][:message.Index], todos[c.ChatID()][message.Index+1:]...)

		return c.Send(fmt.Sprintf(`"%s" has been removed from your todo list`, value))
	}
}
