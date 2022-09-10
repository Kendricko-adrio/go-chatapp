package repository

import (
	"github.com/kendricko-adrio/go-ws/entity"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return GroupRepository{
		db: db,
	}
}

func (group *GroupRepository) GetById(id uint) entity.Group {
	var groupEntity entity.Group
	group.db.Find(&groupEntity, id)
	return groupEntity
}
