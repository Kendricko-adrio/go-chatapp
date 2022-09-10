package entity

import "github.com/gorilla/websocket"

type WSConnect struct {
	User       User
	Connection *websocket.Conn
}
