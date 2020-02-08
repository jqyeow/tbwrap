package main

import (
	"github.com/enrico5b1b4/tbwrap"
)

func HandlePing() func(c tbwrap.Context) error {
	return func(c tbwrap.Context) error {

		return c.Send("pong!")
	}
}
