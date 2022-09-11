package repository

import (
	"log"

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

func (repo *ChatRepository) GetUserChats(userId uint) []entity.GroupDetail {

	var groups []entity.GroupDetail
	var groupId []string
	repo.db.Select("group_id").Find(&groups, "user_id = ?", userId)
	repo.db.Table("users").Select("group_id").Where("user_id = ?", userId).Scan(groupId)
	log.Println(groupId)

	// repo.db.Find(&groups, "group_id IN ? AND user_id != ?", groupId, userId)

	return groups
}
