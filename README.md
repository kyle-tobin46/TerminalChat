# TerminalChat 🖥️💬

A lightweight real-time terminal-based chat application built in Go using WebSockets.

This project includes:
- A Go WebSocket **server** that handles multiple clients
- A **terminal client** that allows users to send and receive messages
- Support for usernames and real-time broadcasting

---

## 📦 Features

- Real-time message broadcasting via WebSockets
- Simple terminal interface for chatting
- Unique usernames with join announcements
- Message echo prevention (clients don't see their own messages twice)
- Deployable locally, over LAN, or with ngrok

---

## 🚀 How to Use

### 🖥️ Run the Server

```bash
go run Server/main.go
```

Then, either run the Client/main.go or execute:
```bash
go run Client/main.go
```
You'll be prompted for a username, then connected to the chat.

## 🌐 Remote Usage with ngrok
To expose your server to the internet:

```bash
ngrok http 8080
```

Then update the client URL in Client/main.go to:
```bash
websocket.DefaultDialer.Dial("wss://your-ngrok-url.ngrok.io/chat", nil)
```
Now others can connect using your public ngrok link.