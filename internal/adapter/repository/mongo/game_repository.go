package mongo

import (
	"github.com/pleum/hexgo/internal/core/domain"
	"github.com/pleum/hexgo/internal/core/port"
)

type GameRepository struct{}

// Create implements port.GameRepo.
func (*GameRepository) Create(entity domain.Game) (int, error) {
	panic("unimplemented")
}

// DeleteByID implements port.GameRepo.
func (*GameRepository) DeleteByID(id int) error {
	panic("unimplemented")
}

// FindAll implements port.GameRepo.
func (*GameRepository) FindAll() ([]*domain.Game, error) {
	panic("unimplemented")
}

// FindByID implements port.GameRepo.
func (*GameRepository) FindByID(id int) (*domain.Game, error) {
	panic("unimplemented")
}

// Update implements port.GameRepo.
func (*GameRepository) Update(entity domain.Game) error {
	panic("unimplemented")
}

func NewMongoGameRepository() port.GameRepo {
	return &GameRepository{}
}
