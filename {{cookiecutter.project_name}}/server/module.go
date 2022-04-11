package server

import (
	"go.uber.org/fx"
	"{{cookiecutter.project_name}}/server/handler"
	"{{cookiecutter.project_name}}/server/hook"
)

var Module = fx.Options(
	fx.Provide(handler.NewHandler),
	fx.Invoke(hook.RegisterHooks),
)
