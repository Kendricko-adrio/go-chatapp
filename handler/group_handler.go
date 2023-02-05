package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kendricko-adrio/go-ws/service/group"
)

type GroupHandler struct {
	groupService group.GroupService
}

func NewGroupHandler(groupService group.GroupService) GroupHandler {
	return GroupHandler{
		groupService: groupService,
	}
}

func (handler *GroupHandler) GetGroupByUser(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	username, _ := variable["username"]
	groupDetails := handler.groupService.GetGroupByUser(username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groupDetails)
}
