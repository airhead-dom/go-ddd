package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ddd/util/logger"
	"net/http"
	"time"
)

type PingHandler struct {
	route  string
	method string
	logger logger.Logger
}

func NewPingHandler(logger logger.Logger) WebHandler {
	return &PingHandler{route: "/ping", method: http.MethodGet, logger: logger}
}

func (h *PingHandler) Route() string {
	return h.route
}

func (h *PingHandler) Method() string {
	return h.method
}

func (h *PingHandler) Handler() gin.HandlerFunc {
	return func(gc *gin.Context) {
		h.logger.Log(fmt.Sprintf("handling ping request at %v", time.Now().UTC().Add(time.Hour*7).Format("02-01-2006 15:04:05")))
		gc.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
