package providers

import (
	controller "go-starter/app/conroller"
	"go-starter/app/repository"
	"go-starter/app/route"
	"go-starter/app/service"

	"go.uber.org/fx"
)

// 提供 Service 层依赖
func forService() fx.Option {
	return fx.Provide(
		service.NewBookService,
	)
}

// 提供 Repository 层依赖
func forRepository() fx.Option {
	return fx.Provide(
		repository.NewBookRepository,
	)
}

// 提供 Controller 层依赖
func forController() fx.Option {
	return fx.Provide(
		controller.NewBookController,
	)
}

// 提供路由处理器依赖
func forRouteHandler() fx.Option {
	return fx.Provide(
		route.NewRouteHandler,
	)
}

func Init() fx.Option {
	return fx.Options(
		forService(),
		forRepository(),
		forController(),
		forRouteHandler(),
	)
}
