// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/pleum/hexgo/cmd/api/internal/app"
	"github.com/pleum/hexgo/internal/adapter/repository"
	"github.com/pleum/hexgo/internal/adapter/repository/mongo"
	"github.com/pleum/hexgo/internal/core/service/game"
)

// Injectors from wire.go:

func InitializeContainer() (Container, func(), error) {
	client, cleanup, err := repository.NewMongoClient()
	if err != nil {
		return Container{}, nil, err
	}
	gameRepo := mongo.NewMongoGameRepository(client)
	playerRepo := mongo.NewMongoPlayerRepositry(client)
	service := game.NewGameService(gameRepo, playerRepo)
	handler := app.NewHandler(service)
	fiberApp, cleanup2, err := app.NewFiberApp(handler)
	if err != nil {
		cleanup()
		return Container{}, nil, err
	}
	container := Container{
		App: fiberApp,
	}
	return container, func() {
		cleanup2()
		cleanup()
	}, nil
}
