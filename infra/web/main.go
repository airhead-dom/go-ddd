package web

import (
	"context"
	"fmt"
	"go-ddd/infra/web/handlers"
	"go-ddd/util/logger"
	"go.uber.org/fx"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGinEngine(handlers []handlers.WebHandler) *gin.Engine {
	r := gin.Default()

	for _, handler := range handlers {
		r.Handle(handler.Method(), handler.Route(), handler.Handler())
	}

	return r
}

func NewHttpServer(lc fx.Lifecycle, ginEngine *gin.Engine, logger logger.Logger) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: ginEngine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			logger.Log(fmt.Sprintf("Starting HTTP server at", srv.Addr))
			go func() {
				_ = srv.Serve(ln)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
