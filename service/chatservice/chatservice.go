package chatservice

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/kendricko-adrio/go-ws/db"
	"github.com/kendricko-adrio/go-ws/entity"
	"github.com/kendricko-adrio/go-ws/repository"
	"github.com/kendricko-adrio/go-ws/service/group"
)

func GetAllChat(userId uint) []entity.GroupDetail {
	repo := repository.NewChatRepository(db.GetDBInstance())

	return repo.GetUserChats(userId)
}

func Receive(ws entity.WSConnect) {
	for {
		message := entity.Message{}
		err := ws.Connection.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			return
		}

		groupService := group.WireGroupService()

		users := groupService.GetToUser(ws.User.Id, message.To)

		log.Println(users)
		if len(users) == 0 {
			continue
		}
		chatRepo := repository.NewChatRepository(db.GetDBInstance())

		chat := chatRepo.InsertChat(message.Message, users[0].Group, ws.User)

		for _, val := range users {
			log.Println(val.User.Username)
			if test, ok := entity.Connections[val.User.Username]; ok {
				Send(test.Connection, chat, 1)
			}
		}
	}
}

func Send(conn *websocket.Conn, message entity.Chat, messageType int) {
	if err := conn.WriteJSON(message); err != nil {
		log.Println(err)
		return
	}
}
