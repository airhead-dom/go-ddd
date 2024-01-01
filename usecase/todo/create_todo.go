package todo

import (
	"fmt"
	"go-ddd/core/domain/todo"
	"go-ddd/infra/db/repository"
	"go-ddd/usecase"
	"go-ddd/util/logger"
)

type CreateTodoDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnedBy     uint   `json:"ownedBy"`
}

type CreateTodoUC struct {
	logger   logger.Logger
	todoRepo *repository.TodoRepository
}

func NewCreateTodoUC(logger logger.Logger, todoRepo *repository.TodoRepository) usecase.UseCase[CreateTodoDto] {
	return &CreateTodoUC{logger: logger, todoRepo: todoRepo}
}

func (c *CreateTodoUC) Do(param CreateTodoDto) (any, error) {
	ent := todo.TodoEntity{
		Title:       param.Title,
		Description: param.Description,
		OwnedBy:     param.OwnedBy,
	}

	created, err := c.todoRepo.Create(ent)
	if err != nil {
		c.logger.Log(fmt.Sprintf("failed creating todo. err: %v", err))
		return nil, err
	}

	c.logger.Log(fmt.Sprintf("successfully created a todo. id: %v", created.Id))

	return created, nil
}
