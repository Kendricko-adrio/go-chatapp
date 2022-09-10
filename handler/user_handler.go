package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kendricko-adrio/go-ws/db"
	"github.com/kendricko-adrio/go-ws/repository"
	"github.com/kendricko-adrio/go-ws/service/user"
)

type UserHandler struct {
	Service user.UserService
}

func GetUserHandlerWired() UserHandler {
	dbConnection := db.GetDBInstance()
	repo := repository.NewRepo(dbConnection)
	userService := user.UserService{Repo: repo}
	return UserHandler{Service: userService}
}

func (handler *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {

	variable := mux.Vars(r)
	id, _ := strconv.ParseUint(variable["id"], 0, 8)
	user := handler.Service.FindById(uint(id))

	json.NewEncoder(w).Encode(user)

}
