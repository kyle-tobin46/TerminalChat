package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
)


/*	 	 *** Global Variables ***		*/

// Upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // allow all connections
}

var clients = make(map[*websocket.Conn]string)

var broadcast = make(chan string)


func main() {
	http.HandleFunc("/chat", wsHandler)
	http.HandleFunc("/clients", clientsHandler)
	fmt.Println("Server started at http://localhost:8080/chat")

	go func() {
		for {
			msg := <-broadcast  
			for client := range clients {
				err := client.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
	}()

	http.ListenAndServe("0.0.0.0:8080", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	_, msg, err:= conn.ReadMessage()
	if err != nil {
		return
	}
	// msg is their username

	clients[conn] = string(msg)

	go func(c *websocket.Conn) {
		defer conn.Close()
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				delete(clients, c)
				c.Close()
				break
			}
			username := clients[c]
			broadcast <- username + ": " + string(msg)
		}
	}(conn)

	
}

func clientsHandler(w http.ResponseWriter, r *http.Request) {
	var list []string
	for conn := range clients {
		list = append(list, conn.RemoteAddr().String())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}