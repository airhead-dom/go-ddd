package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ddd/infra/web/handlers"
	"go-ddd/usecase"
	"go-ddd/util/logger"
	"go.uber.org/fx"
	"net/http"
)

type HdGetUsers struct {
	fx.In

	route         string
	method        string
	logger        logger.Logger
	getAllUsersUc usecase.UseCase[any]
}

type HdGetUsersParams struct {
	fx.In

	Logger logger.Logger
	Uc     usecase.UseCase[any] `name:"getAllUsersUC"`
}

func NewGetUsersHandler(p HdGetUsersParams) handlers.WebHandler {
	return &HdGetUsers{
		route:         "/users",
		method:        http.MethodGet,
		logger:        p.Logger,
		getAllUsersUc: p.Uc,
	}
}

func (h *HdGetUsers) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		result, err := h.getAllUsersUc.Do(nil)
		if err != nil {
			context.JSON(http.StatusBadRequest, handlers.HttpResponse{
				Message: fmt.Sprintf("failed getting user list. err: %v", err),
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

func (h *HdGetUsers) Route() string {
	return h.route
}

func (h *HdGetUsers) Method() string {
	return h.method
}
