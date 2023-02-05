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
	var groupId []int

	repo.db.Table("group_details").Select("group_id").Where("user_id = ?", userId).Scan(&groupId)
	repo.db.Preload("User").Find(&groups, "user_id != ? AND group_id IN ?", userId, groupId)
	log.Println(groupId)
	log.Println("test")

	// repo.db.Find(&groups, "group_id IN ? AND user_id != ?", groupId, userId)

	return groups
}

func (repo *ChatRepository) GetChatsByGroup(groupId uint) []entity.Chat {
	var chats []entity.Chat
	repo.db.Find(&chats, "group_id = ?", groupId)
	return chats
}
