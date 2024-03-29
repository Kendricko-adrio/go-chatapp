package group

import (
	"github.com/kendricko-adrio/go-ws/db"
	"github.com/kendricko-adrio/go-ws/entity"
	"github.com/kendricko-adrio/go-ws/repository"
	"github.com/kendricko-adrio/go-ws/service/user"
)

type GroupService struct {
	groupRepo       repository.GroupRepository
	groupDetailRepo repository.GroupDetailRepository
	userService     user.UserService
}

func NewGroupService(groupRepo repository.GroupRepository, groupDetailRepo repository.GroupDetailRepository) GroupService {
	return GroupService{
		groupRepo:       groupRepo,
		groupDetailRepo: groupDetailRepo,
		userService:     user.UserService{Repo: repository.NewRepo(db.GetDBInstance())},
	}
}

func NewGroupServiceWired() GroupService {
	db := db.GetDBInstance()
	return GroupService{
		groupRepo:       repository.NewGroupRepository(db),
		groupDetailRepo: repository.NewGroupDetailRepository(db),
		userService:     user.UserService{Repo: repository.NewRepo(db)},
	}
}

func WireGroupService() GroupService {
	db := db.GetDBInstance()
	return GroupService{
		groupRepo:       repository.NewGroupRepository(db),
		groupDetailRepo: repository.NewGroupDetailRepository(db),
	}
}

func (service *GroupService) GetGroupByUserAndGroupId(userID uint, groupID uint) entity.Group {

	return service.groupDetailRepo.GetGroupByUserAndGroupId(userID, groupID)
}

func (service *GroupService) GetToUser(userFrom uint, groupID uint) []entity.GroupDetail {

	// group := service.groupDetailRepo.GetGroupByUserAndGroupId(userFrom, groupID)
	group := service.groupDetailRepo.GetUserByGroupIDAndNotUser(groupID, userFrom)
	return group
}

func (service *GroupService) GetGroupByUser(username string) []entity.GroupDetail {

	user := service.userService.FindByUsername(username)
	return service.groupDetailRepo.GetGroupByUser(user.Id)
}
