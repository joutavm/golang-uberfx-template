package main

import (
	"go.uber.org/fx"
	"{{cookiecutter.project_name}}/config"
	"{{cookiecutter.project_name}}/database"
	"{{cookiecutter.project_name}}/product"
	"{{cookiecutter.project_name}}/server"
)

func main() {
	fx.New(
		config.Module,
		database.Module,
		server.Module,
		product.Module,
	).Run()
}
