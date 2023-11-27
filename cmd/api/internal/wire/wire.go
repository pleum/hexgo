//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/pleum/hexgo/cmd/api/internal/app"
	// internalWire "github.com/pleum/hexgo/internal/wire"
)

func InitializeContainer() (Container, func(), error) {
	wire.Build(
		// internalWire.DatabaseConnectionSet,
		// internalWire.MongoRepositorySet,
		app.NewFiberApp,
		wire.Struct(new(Container), "*"),
	)
	return Container{}, func() {}, nil
}
