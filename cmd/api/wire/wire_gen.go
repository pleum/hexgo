// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/pleum/hexgo/cmd/api/internal/adapter"
)

// Injectors from wire.go:

func InitializeEvent() (adapter.Event, error) {
	event := adapter.NewEvent()
	return event, nil
}
