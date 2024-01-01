package user

import "time"

type UserEntity struct {
	Id        uint
	Name      string
	Email     string
	CreatedAt time.Time
}
