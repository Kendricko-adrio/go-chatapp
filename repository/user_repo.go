package repository

import (
	"errors"
	"log"

	"github.com/kendricko-adrio/go-ws/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo *UserRepository) FindAll() []entity.User {
	return nil
}

func (repo *UserRepository) FindById(id uint) entity.User {

	var user entity.User
	// repo.db.First(&user, id)
	query := repo.db.Find(&user, "id = ?", id)
	log.Println(user.GetUsername())
	log.Println(query.RowsAffected)
	return user
}

func (repo *UserRepository) FindByUsername(username string) entity.User {

	var user entity.User
	// repo.db.First(&user, id)
	query := repo.db.Find(&user, "username = ?", username)
	log.Println(user)
	log.Println(query.RowsAffected)
	return user
}

func (repo *UserRepository) FindByUsernameAndPassword(username string, password string) (entity.User, error) {

	var user entity.User
	query := repo.db.Find(&user, "username = ? AND password = ?", username, password)
	log.Println("query: ", query.RowsAffected)
	if query.RowsAffected == 0 {
		return user, errors.New("User not found")
	}
	return user, nil
}
