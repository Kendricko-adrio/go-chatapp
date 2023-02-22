package entity

import "time"

type Session struct {
	User    User
	Expired time.Time
}
