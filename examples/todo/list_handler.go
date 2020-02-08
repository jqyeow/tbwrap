package main

import (
	"bytes"
	"html/template"

	"github.com/enrico5b1b4/tbwrap"
)

func HandleList(todos map[int][]string) func(c tbwrap.Context) error {
	return func(c tbwrap.Context) error {
		chatTodos := todos[c.ChatID()]
		if len(chatTodos) == 0 {
			return c.Send("your todo list is empty")
		}

		t := template.Must(template.New("text").Parse(text))
		var buf bytes.Buffer
		if err := t.Execute(&buf, chatTodos); err != nil {
			return err
		}

		return c.Send(buf.String())
	}
}

const text = `{{ range $i, $entry := . }}{{printf "%d - %s\n" $i $entry}}{{ end }}`
