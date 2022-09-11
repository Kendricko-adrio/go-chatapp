package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbConnect *gorm.DB = nil

func GetDBInstance() *gorm.DB {
	if dbConnect == nil {
		dbConnect = connect()
	}
	return dbConnect
}

func connect() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSL_MODE")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=" + sslMode + " TimeZone=Asia/Bangkok"

	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
		},
	), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&entity.User{})

	return db
}
