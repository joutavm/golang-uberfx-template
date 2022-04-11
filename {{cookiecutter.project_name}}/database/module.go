package database

import (
	"go.uber.org/fx"
	"{{cookiecutter.project_name}}/database/base"
)

var Module = fx.Options(fx.Provide(base.Connect))
