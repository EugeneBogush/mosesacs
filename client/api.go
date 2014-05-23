package client

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"os"
//	"strings"
)

func Connect(url string) {
	origin := "http://localhost/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println("Error connecting to remote MosesACS instance")
		os.Exit(1)
	}

	var channel = make(chan string)
	go Write(ws, channel)
	go Read(ws, channel)

	for {
		cmd := <-channel

		fmt.Printf("Got '%s' from channel\n", cmd)

		switch {
		  case cmd == "quit":
			  fmt.Println("Quit")
			  os.Exit(0)
		}

	}

}

func Write(ws *websocket.Conn, channel chan string) {
	if _, err := ws.Write([]byte("list\n")); err != nil {
		channel <- "quit"
	}
}

func Read(ws *websocket.Conn, channel chan string) {
	var msg = make([]byte, 512)
	for {
		if _, err := ws.Read(msg); err != nil {
			channel <- "quit"
		}
		channel <- string(msg)
	}

}