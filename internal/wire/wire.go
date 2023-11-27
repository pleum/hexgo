package wire

import (
	"github.com/google/wire"
	"github.com/pleum/hexgo/internal/adapter/repository"
	"github.com/pleum/hexgo/internal/adapter/repository/mongo"
)

var DatabaseConnectionSet = wire.NewSet(
	repository.NewMongoConnection,
)

var MongoRepositorySet = wire.NewSet(
	mongo.NewMongoGameRepository,
	mongo.NewMongoPlayerRepositry,
)
