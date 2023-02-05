package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kendricko-adrio/go-ws/db"
	"github.com/kendricko-adrio/go-ws/entity"
	"github.com/kendricko-adrio/go-ws/repository"
	"github.com/kendricko-adrio/go-ws/service/chatservice"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		panic("error connection")
	}

	username := r.URL.Query().Get("username")
	// asdf
	to := r.URL.Query().Get("to")

	log.Printf("Username: %s, to user : %s", username, to)

	userRepo := repository.NewRepo(db.GetDBInstance())

	user := userRepo.FindByUsername(username)
	// user := entity.NewUser(username, conn, to)
	wsConnection := entity.WSConnect{
		User:       user,
		Connection: conn,
	}
	entity.Connections[username] = wsConnection
	go chatservice.Receive(wsConnection)
}
