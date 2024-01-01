package handlers

import "github.com/gin-gonic/gin"

type WebHandler interface {
	Handler() gin.HandlerFunc
	Route() string
	Method() string
}
