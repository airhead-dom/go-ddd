package di

import (
	"fmt"
	"go-ddd/infra/db"
	"go-ddd/infra/db/repository"
	"go-ddd/infra/web"
	"go-ddd/infra/web/handlers"
	todoHandlers "go-ddd/infra/web/handlers/todo"
	userHandlers "go-ddd/infra/web/handlers/user"
	"go-ddd/usecase/todo"
	user "go-ddd/usecase/user"
	"go-ddd/util/logger"
	"go-ddd/util/mapper"
	"go.uber.org/fx"
	"net/http"
)

func Init() *fx.App {
	return fx.New(
		fx.Provide(
			web.NewHttpServer,
			fx.Annotate(
				web.NewGinEngine,
				fx.ParamTags(`group:"handlers"`),
			),
			AsWebHandler(handlers.NewPingHandler),
			AsWebHandler(todoHandlers.NewCreateTodoHandler),
			AsWebHandler(userHandlers.NewCreateUserHandler),
			AsWebHandler(userHandlers.NewGetUsersHandler),
			AsUseCase(todo.NewCreateTodoUC, "createTodoUC"),
			AsUseCase(user.NewCreateUserUC, "createUserUC"),
			AsUseCase(user.NewGetAllUsersUc, "getAllUsersUC"),
			repository.NewTodoRepository,
			repository.NewUserRepository,
			db.DB,
			mapper.NewGoModelMapper,
			logger.NewZapLogger,
			logger.GetZap,
		),
		//fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		//	return &fxevent.ZapLogger{Logger: log}
		//}),
		fx.Invoke(func(srv *http.Server) {}),
	)
}

func AsWebHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(handlers.WebHandler)),
		fx.ResultTags(`group:"handlers"`),
	)
}

func AsUseCase(f any, name string) any {
	return fx.Annotate(
		f,
		fx.ResultTags(fmt.Sprintf(`name:"%s"`, name)),
	)
}
