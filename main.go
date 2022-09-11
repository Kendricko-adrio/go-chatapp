package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/kendricko-adrio/go-ws/db"
	"github.com/kendricko-adrio/go-ws/entity"
	"github.com/kendricko-adrio/go-ws/handler"
	"github.com/kendricko-adrio/go-ws/repository"
	"github.com/kendricko-adrio/go-ws/service/chatservice"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic("error connection")
	}

	username := r.URL.Query().Get("username")

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

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("html/index.html")
	if err != nil {
		panic(err)
	}
	template.Execute(w, nil)
}

func main() {
	// handler.LoadDotEnvFile()
	loadErr := godotenv.Load(".env")
	if loadErr != nil {
		log.Println("Something wrong with load .env file")
		os.Exit(1)
		return
	}
	router := mux.NewRouter()
	// db.MigrateDB()
	//wiring
	userHandler := handler.GetUserHandlerWired()

	router.HandleFunc("/ws", websocketHandler)
	router.HandleFunc("/", htmlHandler)
	router.HandleFunc("/chat/{userId}", handler.GetUserChats).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", userHandler.GetUserById).Methods(http.MethodGet)
	log.Println("run on port 8080")
	http.ListenAndServe(":8080", router)

}
