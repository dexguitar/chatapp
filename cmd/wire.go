//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dexguitar/chatapp/configs"
	"github.com/dexguitar/chatapp/db"
	"github.com/dexguitar/chatapp/internal/handler"
	"github.com/dexguitar/chatapp/internal/queue"
	"github.com/dexguitar/chatapp/internal/repo"
	"github.com/dexguitar/chatapp/internal/service"
	"github.com/google/wire"
)

func InitApplication() (*application, error) {
	wire.Build(
		configs.LoadConfig,
		repo.NewUserRepo,
		db.NewPostgresDB,
		queue.NewKafkaProducer,
		queue.NewKafkaConsumer,
		queue.New,
		service.NewUserService,
		wire.Bind(new(service.Repo), new(*repo.UserRepository)),
		handler.NewUserHandler,
		wire.Bind(new(handler.UserService), new(*service.UserService)),
		service.NewMessageService,
		wire.Bind(new(service.Queue), new(*queue.Queue)),
		handler.NewMessageHandler,
		wire.Bind(new(handler.MessageService), new(*service.MessageService)),
		handler.InitRouter,
		newApplication,
		handler.NewWSHandler,
		wire.Bind(new(queue.Hub), new(*service.Hub)),
		service.NewHub,
	)

	return &application{}, nil
}
