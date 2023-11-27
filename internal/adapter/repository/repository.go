package repository

import (
	"github.com/google/wire"
	"github.com/pleum/hexgo/internal/adapter/repository/mongo"
)

var MongoRepoSet = wire.NewSet(mongo.NewMongoGameRepository, mongo.NewMongoPlayerRepositry)
