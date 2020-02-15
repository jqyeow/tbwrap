package main

import (
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/enrico5b1b4/tbwrap"
	tb "gopkg.in/tucnak/telebot.v2"
)

type TeleBot struct {
	done     chan interface{}
	listener net.Listener
	handler  interface{}
}

func (t *TeleBot) Handle(endpoint interface{}, handler interface{}) {
	log.Println(endpoint)
	log.Println(handler)
	log.Println(reflect.TypeOf(handler))
	t.handler = handler
}

func (t *TeleBot) Respond(callback *tb.Callback, responseOptional ...*tb.CallbackResponse) error {
	return nil
}

func (t *TeleBot) Send(to tb.Recipient, what interface{}, options ...interface{}) (*tb.Message, error) {
	return nil, nil
}

func (t *TeleBot) Start() {
	log.Println("Start")
	t.done = make(chan interface{})
	// var err error

	// t.listener, err = net.Listen("tcp", ":8081")
	// if err != nil {
	// 	panic(err)
	// }

	// go func(done chan interface{}, ln net.Listener) {
	// 	ln, err := net.Listen("tcp", ":8081")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println("listening on 8081")
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println("accepted on 8081")

	// 	defer fmt.Println("doWork exited.")
	// 	for {
	// 		select {
	// 		// case s := <-strings:
	// 		//     // Do something interesting
	// 		//     fmt.Println(s)
	// 		case <-done:
	// 			return
	// 		default:
	// 			log.Println("in default")
	// 			// will listen for message to process ending in newline (\n)
	// 			message, _ := bufio.NewReader(conn).ReadString('\n')
	// 			// output message received
	// 			fmt.Print("Message Received:", string(message))
	// 			// sample process for string received
	// 			newmessage := strings.ToUpper(message)
	// 			// send new string back to client
	// 			conn.Write([]byte(newmessage + "\n"))
	// 		}
	// 	}
	// }(t.done, t.listener)

	// go func() {
	// 	for {
	// 		// will listen for message to process ending in newline (\n)
	// 		message, _ := bufio.NewReader(conn).ReadString('\n')
	// 		// output message received
	// 		fmt.Print("Message Received:", string(message))
	// 		// sample process for string received
	// 		newmessage := strings.ToUpper(message)
	// 		// send new string back to client
	// 		conn.Write([]byte(newmessage + "\n"))
	// 	}
	// }()
}

func (t *TeleBot) Run() {
	log.Println("run")
	handle, ok := t.handler.(func(m *tb.Message))
	if ok {
		log.Println("ok")
		handle(&tb.Message{Chat: &tb.Chat{}, Text: "ciao"})
		return
	}
	log.Println("not ok")
}

func (t *TeleBot) Stop() {
	close(t.done)
	err := t.listener.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	testTeleBot := &TeleBot{}
	testTeleBot.Start()

	// time.Sleep(5 * time.Second)
	// connect to this socket
	// conn, err := net.Dial("tcp", "127.0.0.1:8081")
	// if err != nil {
	// 	panic(err)
	// }

	called := false
	testTeleBot.Handle("test", func(c tbwrap.Context) error {
		called = true
		return c.Send("test")
	})
	testTeleBot.Run()
	fmt.Println(called)

	// go func(t *TeleBot) {
	// 	time.Sleep(5 * time.Second)
	// 	log.Println("shutting down tcp server")
	// 	testTeleBot.Stop()
	// }(testTeleBot)

	// for {
	// 	// read in input from stdin
	// 	fmt.Println("please enter text")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Text to send: ")
	// 	text, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// send to socket
	// 	fmt.Fprintf(conn, text+"\n")
	// 	// listen for reply
	// 	message, err := bufio.NewReader(conn).ReadString('\n')
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Print("Message from server: " + message)
	// }

}
