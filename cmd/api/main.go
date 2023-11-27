package main

import (
	"log/slog"

	"github.com/pleum/hexgo/cmd/api/wire"
)

func main() {
	_, err := wire.InitializeEvent()
	if err != nil {
		slog.Error(err.Error())
	}
}
