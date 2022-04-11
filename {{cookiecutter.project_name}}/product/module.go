package product

import (
	"go.uber.org/fx"
	"{{cookiecutter.project_name}}/product/controller"
	"{{cookiecutter.project_name}}/product/handler"
	"{{cookiecutter.project_name}}/product/migrate"
	"{{cookiecutter.project_name}}/product/provider"
)

var Module = fx.Options(
	migrate.Module,
	fx.Provide(
		provider.Init,
		handler.Init,
	),
	fx.Invoke(controller.RegisterHandlers),
)
