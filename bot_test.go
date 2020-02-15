package tbwrap_test

import "testing"

func Test_TBWrapBot_Add(t *testing.T) {

	// testTeleBot := &test.TeleBot{}
	// testTeleBot.Start()

	// // connect to this socket
	// conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	// for {
	// 	// read in input from stdin
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Text to send: ")
	// 	text, _ := reader.ReadString('\n')
	// 	// send to socket
	// 	fmt.Fprintf(conn, text+"\n")
	// 	// listen for reply
	// 	message, _ := bufio.NewReader(conn).ReadString('\n')
	// 	fmt.Print("Message from server: " + message)
	// }

	// assert.True(t, true)

	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// mockTeleBot := mocks.NewMockTeleBot(mockCtrl)
	// mockTeleBot.EXPECT().Start().Return()
	// mockTeleBot.EXPECT().Handle(gomock.Any(), gomock.Any()).Return().AnyTimes()
	// testBot := &tbwrap.TBWrapBot{
	// 	TBot: mockTeleBot,
	// }

	// called := false
	// testBot.Add("test", func(c tbwrap.Context) error {
	// 	called = true
	// 	return c.Send("test")
	// })
	// testBot.Start()

	// assert.True(t, called)
}
