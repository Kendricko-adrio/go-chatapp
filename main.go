package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kendricko-adrio/go-ws/handler"
	"github.com/kendricko-adrio/go-ws/service/group"
	"github.com/rs/cors"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func websocketHandler(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		panic("error connection")
// 	}

// 	username := r.URL.Query().Get("username")

// 	to := r.URL.Query().Get("to")

// 	log.Printf("Username: %s, to user : %s", username, to)

// 	userRepo := repository.NewRepo(db.GetDBInstance())

// 	user := userRepo.FindByUsername(username)
// 	// user := entity.NewUser(username, conn, to)
// 	wsConnection := entity.WSConnect{
// 		User:       user,
// 		Connection: conn,
// 	}
// 	entity.Connections[username] = wsConnection
// 	go chatservice.Receive(wsConnection)
// }

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

	groupHandler := handler.NewGroupHandler(group.NewGroupServiceWired())

	router.HandleFunc("/ws", handler.WebsocketHandler)
	router.HandleFunc("/chat/{userId}", handler.GetUserChats).Methods(http.MethodGet)
	router.HandleFunc("/chat/group/{groupId}", handler.GetChatsByGroup).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", userHandler.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/user/username/{username}", userHandler.GetUserByUsername).Methods(http.MethodGet)
	router.HandleFunc("/group-detail/user/{username}", groupHandler.GetGroupByUser).Methods(http.MethodGet)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})
	handle := c.Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	log.Println("run on port " + port)
	http.ListenAndServe(":"+port, handle)

}
