package repository

import (
	"log"

	"github.com/kendricko-adrio/go-ws/entity"
	"gorm.io/gorm"
)

type GroupDetailRepository struct {
	db *gorm.DB
}

func NewGroupDetailRepository(db *gorm.DB) GroupDetailRepository {
	return GroupDetailRepository{
		db: db,
	}
}

func (repo *GroupDetailRepository) GetGroupByUserAndGroupId(userID uint, groupID uint) entity.Group {
	var detail entity.GroupDetail
	repo.db.First(&detail, "user_id = ? AND group_id", userID, groupID)

	return detail.Group
}

func (repo *GroupDetailRepository) GetUserByGroupIDAndNotUser(groupID uint, notUser uint) []entity.GroupDetail {
	var detail []entity.GroupDetail
	repo.db.Preload("User").Preload("Group").Find(&detail, "user_id != ? AND group_id = ?", notUser, groupID)
	return detail
}

func (repo *GroupDetailRepository) GetGroupByUser(userID uint) []entity.GroupDetail {
	var detail []entity.GroupDetail
	repo.db.Preload("User").Preload("Group").Find(&detail, "user_id != ? AND group_id IN (select group_id from group_details where user_id = ?)", userID, userID)
	log.Println(detail)
	return detail
}
