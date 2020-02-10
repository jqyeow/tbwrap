package main

import (
	"fmt"

	"github.com/enrico5b1b4/tbwrap"
)

type HelloMessage struct {
	Name string `regexpGroup:"name"`
}

func HandleHello() func(c tbwrap.Context) error {
	return func(c tbwrap.Context) error {
		message := new(HelloMessage)
		if err := c.Bind(message); err != nil {
			return err
		}

		return c.Send(fmt.Sprintf("Hello %s!", message.Name))
	}
}
