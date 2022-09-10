package db

import "github.com/kendricko-adrio/go-ws/entity"

func MigrateDB() {

	db := GetDBInstance()

	db.Migrator().DropTable(&entity.User{}, &entity.GroupDetail{}, &entity.Group{}, &entity.Chat{})
	db.Migrator().AutoMigrate(&entity.User{}, &entity.GroupDetail{}, &entity.Group{}, &entity.Chat{})
	userSeeder()
	groupSeeder()
	groupDetailSeeder()
}

func userSeeder() {
	db := GetDBInstance()
	db.Create(&entity.User{
		Id:       1,
		Username: "ricko1",
		Password: "password",
		Email:    "email",
	})

	db.Create(&entity.User{
		Id:       2,
		Username: "ricko2",
		Password: "password",
		Email:    "email",
	})
	db.Create(&entity.User{
		Id:       3,
		Username: "ricko3",
		Password: "password",
		Email:    "email",
	})
}

func groupSeeder() {
	db := GetDBInstance()

	db.Create(&entity.Group{
		GroupName: "group1",
		GroupType: "Personal",
	})
}

func groupDetailSeeder() {
	db := GetDBInstance()

	db.Create(&entity.GroupDetail{
		UserID:  1,
		GroupID: 1,
	})
	db.Create(&entity.GroupDetail{
		UserID:  2,
		GroupID: 1,
	})

}
