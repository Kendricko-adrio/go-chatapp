package entity

type User struct {
	Id       uint `gorm:"primaryKey"`
	Username string
	Password string `json:"-"`
	Email    string
}

func NewUser(id uint, username string, password string, email string) User {
	return User{
		Id:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
}

func (user *User) GetUsername() string {
	return user.Username
}

// func NewUser(username string, connection *websocket.Conn, toUser string) *User {
// 	return &User{
// 		username:   username,
// 		connection: connection,
// 		toUser:     toUser,
// 	}
// }

// func (user *User) GetUsername() string {
// 	return user.username
// }

// func (user *User) GetConnection() *websocket.Conn {
// 	return user.connection
// }

// func (user *User) GetToUser() string {
// 	return user.toUser
// }
