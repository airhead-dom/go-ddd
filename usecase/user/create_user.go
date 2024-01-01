package todo

import (
	"fmt"
	"go-ddd/core/domain/user"
	"go-ddd/infra/db/repository"
	"go-ddd/usecase"
	"go-ddd/util/logger"
)

type CreateUserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserUC struct {
	logger   logger.Logger
	userRepo *repository.UserRepository
}

func NewCreateUserUC(logger logger.Logger, userRepo *repository.UserRepository) usecase.UseCase[CreateUserDto] {
	return &CreateUserUC{logger: logger, userRepo: userRepo}
}

func (c *CreateUserUC) Do(param CreateUserDto) (any, error) {
	ent := user.UserEntity{
		Name:  param.Name,
		Email: param.Email,
	}

	created, err := c.userRepo.Create(ent)
	if err != nil {
		c.logger.Log(fmt.Sprintf("failed creating user. err: %v", err))
		return nil, err
	}

	c.logger.Log(fmt.Sprintf("successfully created a user. id: %v", created.Id))

	return created, nil
}
