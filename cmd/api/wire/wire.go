//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/pleum/hexgo/cmd/api/internal/adapter"
)

func InitializeEvent() (adapter.Event, error) {
	wire.Build(adapter.NewEvent)
	return adapter.Event{}, nil
}
