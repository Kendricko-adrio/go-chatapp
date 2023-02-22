package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/kendricko-adrio/go-ws/db"
	"github.com/kendricko-adrio/go-ws/entity"
	"github.com/kendricko-adrio/go-ws/repository"
	"github.com/kendricko-adrio/go-ws/service/user"
	"github.com/kendricko-adrio/go-ws/session"
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (handler *UserHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {

	variable := mux.Vars(r)
	username := variable["username"]
	user := handler.Service.FindByUsername(username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (handler *UserHandler) PostUserIsAuth(w http.ResponseWriter, r *http.Request) {

	user, err := CheckCookie(w, r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&entity.Response{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&entity.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    user,
	})
}

func (handler *UserHandler) PostUserByUsernameAndPassword(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("username: ", username)
	log.Println("password: ", password)
	user, err := handler.Service.FindByUsernameAndPassword(username, password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	cookieUUID := uuid.New().String()
	log.Println("cookieUUID: ", cookieUUID)
	cookie := &http.Cookie{
		Name:     "session",
		Value:    cookieUUID,
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
		Domain:   "localhost",
	}
	// w.Header().Set("Set-Cookie", cookie.String())
	http.SetCookie(w, cookie)
	session.InsertSession(cookieUUID, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
