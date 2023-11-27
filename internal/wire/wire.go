package wire

import (
	"github.com/google/wire"
	"github.com/pleum/hexgo/internal/adapter/repository"
	"github.com/pleum/hexgo/internal/adapter/repository/mongo"
	"github.com/pleum/hexgo/internal/core/service/game"
)

var DatabaseClientSet = wire.NewSet(
	repository.NewMongoClient,
)

var MongoRepositorySet = wire.NewSet(
	mongo.NewMongoGameRepository,
	mongo.NewMongoPlayerRepositry,
)

var ServiceSet = wire.NewSet(game.NewGameService)
