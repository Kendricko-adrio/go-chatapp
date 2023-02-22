package user

import (
	"github.com/kendricko-adrio/go-ws/entity"
	"github.com/kendricko-adrio/go-ws/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (user *UserService) FindAll() []entity.User {
	return nil
}

func (user *UserService) FindByUsername(username string) entity.User {
	return user.Repo.FindByUsername(username)
}

func (service *UserService) FindById(id uint) entity.User {

	return service.Repo.FindById(id)
}

func (user *UserService) FindByUsernameAndPassword(username string, password string) (entity.User, error) {
	return user.Repo.FindByUsernameAndPassword(username, password)
}
