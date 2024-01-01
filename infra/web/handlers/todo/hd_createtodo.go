package todo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ddd/infra/web/handlers"
	"go-ddd/usecase"
	"go-ddd/usecase/todo"
	"go-ddd/util/logger"
	"go.uber.org/fx"
	"net/http"
)

type HdCreateTodo struct {
	route        string
	method       string
	logger       logger.Logger
	createTodoUc usecase.UseCase[todo.CreateTodoDto]
}

type HdCreateTodoParams struct {
	fx.In

	Logger logger.Logger
	Uc     usecase.UseCase[todo.CreateTodoDto] `name:"createTodoUC"`
}

func NewCreateTodoHandler(p HdCreateTodoParams) handlers.WebHandler {
	return &HdCreateTodo{
		route:        "/todos",
		method:       http.MethodPost,
		logger:       p.Logger,
		createTodoUc: p.Uc,
	}
}

func (h *HdCreateTodo) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		body := todo.CreateTodoDto{}
		err := context.ShouldBindJSON(&body)
		if err != nil {
			h.logger.Log(fmt.Sprintf("failed binding request body. err: %v", err))
			context.JSON(http.StatusBadRequest, handlers.HttpResponse{
				Message: fmt.Sprintf("failed binding request body. err: %v", err),
				Code:    "99",
				Data:    nil,
			})
			return
		}

		result, err := h.createTodoUc.Do(body)
		if err != nil {
			h.logger.Log(fmt.Sprintf("failed creating todo. err: %v", err))

			context.JSON(http.StatusBadRequest, handlers.HttpResponse{
				Message: fmt.Sprintf("failed creating todo. err: %v", err),
				Code:    "99",
				Data:    nil,
			})

			return
		}

		context.JSON(http.StatusOK, handlers.HttpResponse{
			Message: "ok",
			Code:    "00",
			Data:    result,
		})
	}
}

func (h *HdCreateTodo) Route() string {
	return h.route
}

func (h *HdCreateTodo) Method() string {
	return h.method
}
