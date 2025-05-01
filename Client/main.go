package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/gorilla/websocket"
)

func main() {
	fmt.Print("Enter your username: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	username := input.Text()

	conn, _, err := websocket.DefaultDialer.Dial("ws://0.0.0.0:8080/chat", nil)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	conn.WriteMessage(websocket.TextMessage, []byte(username))
	conn.WriteMessage(websocket.TextMessage, []byte("🟢 " + username + " has joined the chat"))
	fmt.Println("Connected to chat server as", username)

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Read error:", err)
				return
			}
			text := string(msg)

			if strings.HasPrefix(text, username+":") {
				continue
			}

			fmt.Println(text)
		}
	}()

	for input.Scan() {
		text := input.Text()
		conn.WriteMessage(websocket.TextMessage, []byte(text))
	}
}
