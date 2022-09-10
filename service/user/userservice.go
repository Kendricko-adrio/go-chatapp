package user

import (
	"github.com/kendricko-adrio/go-ws/entity"
	"github.com/kendricko-adrio/go-ws/repository"
)

type UserService struct {
	Repo repository.IRepository[entity.User, uint]
}

func (user *UserService) FindAll() []entity.User {
	return nil
}

func (service *UserService) FindById(id uint) entity.User {

	return service.Repo.FindById(id)
}
