package entity

type Message struct {
	To      uint   `json:"to"` // groupId
	Message string `json:"message"`
}
