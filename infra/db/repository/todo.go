package repository

import (
	"database/sql"
	"go-ddd/core/domain/todo"
	"go-ddd/util/mapper"
	"gorm.io/gorm"
	"time"
)

type TodoEntityModel struct {
	Id          uint
	Title       string
	Description sql.NullString
	CreatedAt   time.Time     `gorm:"column:created_at"`
	CreatedBy   sql.NullInt64 `gorm:"column:created_by"`
	OwnedBy     sql.NullInt64 `gorm:"column:owned_by"`
	Order       uint
}

func (_ TodoEntityModel) TableName() string {
	return "todos"
}

type TodoRepository struct {
	db     *gorm.DB
	mapper mapper.Mapper
}

func NewTodoRepository(db *gorm.DB, mapper mapper.Mapper) *TodoRepository {
	return &TodoRepository{db: db, mapper: mapper}
}

func (r *TodoRepository) Create(model todo.TodoEntity) (todo.TodoEntity, error) {
	em := TodoEntityModel{
		Title: model.Title,
		Description: sql.NullString{
			String: model.Description,
			Valid:  true,
		},
		CreatedBy: sql.NullInt64{Int64: int64(model.OwnedBy), Valid: true},
		OwnedBy:   sql.NullInt64{Int64: int64(model.OwnedBy), Valid: true},
		Order:     1,
	}

	ent := todo.TodoEntity{}

	res := r.db.Create(&em)

	if res.Error != nil {
		return ent, res.Error
	}

	r.mapModel(em, &ent)

	return ent, nil
}

func (r *TodoRepository) mapModel(model TodoEntityModel, ent *todo.TodoEntity) {
	r.mapper.Map(ent, model)
	ent.Description = model.Description.String
	ent.CreatedBy = uint(model.CreatedBy.Int64)
	ent.OwnedBy = uint(model.OwnedBy.Int64)
}
