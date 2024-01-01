package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ddd/infra/web/handlers"
	"go-ddd/usecase"
	user "go-ddd/usecase/user"
	"go-ddd/util/logger"
	"go.uber.org/fx"
	"net/http"
)

type HdCreateUser struct {
	route        string
	method       string
	logger       logger.Logger
	createUserUc usecase.UseCase[user.CreateUserDto]
}

type HdCreateUserParams struct {
	fx.In

	Logger logger.Logger
	Uc     usecase.UseCase[user.CreateUserDto] `name:"createUserUC"`
}

func NewCreateUserHandler(p HdCreateUserParams) handlers.WebHandler {
	return &HdCreateUser{
		route:        "/users",
		method:       http.MethodPost,
		logger:       p.Logger,
		createUserUc: p.Uc,
	}
}

func (h *HdCreateUser) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		body := user.CreateUserDto{}
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

		result, err := h.createUserUc.Do(body)
		if err != nil {
			context.JSON(http.StatusBadRequest, handlers.HttpResponse{
				Message: fmt.Sprintf("failed creating user. err: %v", err),
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

func (h *HdCreateUser) Route() string {
	return h.route
}

func (h *HdCreateUser) Method() string {
	return h.method
}
