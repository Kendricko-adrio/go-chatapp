package repository

import (
	"github.com/kendricko-adrio/go-ws/entity"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return ChatRepository{
		db: db,
	}
}

func (repo *ChatRepository) InsertChat(message string, group entity.Group, user entity.User) entity.Chat {
	chat := entity.Chat{
		Message:  message,
		Group:    group,
		UserFrom: user,
	}
	repo.db.Create(&chat)
	return chat
}
