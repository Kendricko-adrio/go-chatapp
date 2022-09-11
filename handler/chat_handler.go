package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kendricko-adrio/go-ws/service/chatservice"
)

func GetUserChats(w http.ResponseWriter, r *http.Request) {

	variable := mux.Vars(r)
	temp, _ := strconv.ParseUint(variable["userId"], 0, 8)
	id := uint(temp)

	chats := chatservice.GetAllChat(id)

	json.NewEncoder(w).Encode(chats)
}
