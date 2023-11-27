package game

import (
	"github.com/pleum/hexgo/internal/core/domain"
	"github.com/pleum/hexgo/internal/core/port"
)

type Service struct {
	GameRepo   port.GameRepo
	PlayerRepo port.PlayerRepo
}

func (s Service) CreateNew(players []domain.Player) (*domain.Game, error) {
	g := new(domain.Game)

	returnedGame, err := s.GameRepo.Create(*g)
	if err != nil {
		return nil, err
	}

	returnedPlayers, err := s.PlayerRepo.CreateMany(players)
	if err != nil {
		return nil, err
	}

	returnedGame.Players = returnedPlayers

	return &returnedGame, nil
}

func NewGameService(
	gameRepo port.GameRepo,
	playerRepo port.PlayerRepo,
) Service {
	return Service{
		GameRepo:   gameRepo,
		PlayerRepo: playerRepo,
	}
}
