package session

import (
	"log"
	"os"
	"time"

	"github.com/kendricko-adrio/go-ws/entity"
)

// uuid -> user
var session map[string]entity.Session = make(map[string]entity.Session)

func InsertSession(username string, user entity.User) {
	second := os.Getenv("SESSION_EXPIRED_SECOND")
	secondInt, _ := time.ParseDuration(second + "m")

	session[username] = entity.Session{
		User:    user,
		Expired: time.Now().Add(secondInt),
	}
	log.Println("session: ", session)
}

func DeleteUserSession(uuid string) {
	delete(session, uuid)
}

func GetUserSession(uuid string) (entity.User, bool) {
	value, ok := session[uuid]
	return value.User, ok
}

func GetSession(uuid string) (entity.Session, bool) {
	value, ok := session[uuid]
	return value, ok
}
