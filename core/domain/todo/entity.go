package todo

import "time"

type TodoEntity struct {
	Id          uint
	Title       string
	Description string
	CreatedAt   time.Time
	CreatedBy   uint
	OwnedBy     uint
	Order       uint
}
