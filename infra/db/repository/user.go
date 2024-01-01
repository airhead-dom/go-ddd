package repository

import (
	"go-ddd/core/domain/user"
	"go-ddd/util/mapper"
	"gorm.io/gorm"
	"time"
)

type UserEntityModel struct {
	Id        uint
	Name      string
	Email     string
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (_ UserEntityModel) TableName() string {
	return "users"
}

type UserRepository struct {
	db     *gorm.DB
	mapper mapper.Mapper
}

func NewUserRepository(db *gorm.DB, mapper mapper.Mapper) *UserRepository {
	return &UserRepository{db: db, mapper: mapper}
}

func (r *UserRepository) List() ([]user.UserEntity, error) {
	var users []UserEntityModel
	var entities []user.UserEntity

	res := r.db.Find(&users)
	if res.Error != nil {
		return entities, res.Error
	}

	for _, usr := range users {
		var ent user.UserEntity
		_ = r.mapper.Map(&ent, usr)

		entities = append(entities, ent)
	}

	return entities, nil
}

func (r *UserRepository) Create(model user.UserEntity) (user.UserEntity, error) {
	em := UserEntityModel{
		Name:  model.Name,
		Email: model.Email,
	}

	ent := user.UserEntity{}

	res := r.db.Create(&em)

	if res.Error != nil {
		return ent, res.Error
	}

	r.mapper.Map(&ent, em)

	return ent, nil
}
