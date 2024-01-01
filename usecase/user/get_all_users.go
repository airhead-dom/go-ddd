package todo

import (
	"fmt"
	"go-ddd/infra/db/repository"
	"go-ddd/usecase"
	"go-ddd/util/logger"
	"go-ddd/util/mapper"
)

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetAllUsersUc struct {
	logger   logger.Logger
	userRepo *repository.UserRepository
	mapper   mapper.Mapper
}

func NewGetAllUsersUc(logger logger.Logger, userRepo *repository.UserRepository, mapper mapper.Mapper) usecase.UseCase[any] {
	return &GetAllUsersUc{logger: logger, userRepo: userRepo, mapper: mapper}
}

func (uc *GetAllUsersUc) Do(_ any) (any, error) {
	list, err := uc.userRepo.List()
	if err != nil {
		uc.logger.Log(fmt.Sprintf("failed getting users list. err: %v", err))
		return nil, err
	}

	var dtos []UserDto
	for _, entity := range list {
		var dto UserDto
		_ = uc.mapper.Map(&dto, entity)

		dtos = append(dtos, dto)
	}

	return dtos, nil
}
