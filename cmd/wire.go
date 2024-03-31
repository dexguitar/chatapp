//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dexguitar/chatapp/configs"
	"github.com/dexguitar/chatapp/db"
	"github.com/dexguitar/chatapp/internal/handler"
	"github.com/dexguitar/chatapp/internal/repo"
	"github.com/dexguitar/chatapp/internal/service"
	"github.com/google/wire"
)

func InitApplication() (*application, error) {
	wire.Build(
		configs.LoadConfig,
		repo.NewUserRepo,
		db.NewPostgresDB,
		service.NewUserService,
		wire.Bind(new(service.Repo), new(*repo.UserRepository)),
		handler.NewUserHandler,
		wire.Bind(new(handler.UserService), new(*service.UserService)),
		handler.InitRouter,
		newApplication,
	)

	return &application{}, nil
}