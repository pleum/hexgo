package mongo

import (
	"github.com/pleum/hexgo/internal/core/domain"
	"github.com/pleum/hexgo/internal/core/port"
)

type GameRepository struct{}

// Create implements port.PlayerRepo.
func (*GameRepository) Create(entity domain.Player) (int, error) {
	panic("unimplemented")
}

// DeleteByID implements port.PlayerRepo.
func (*GameRepository) DeleteByID(id int) error {
	panic("unimplemented")
}

// FindAll implements port.PlayerRepo.
func (*GameRepository) FindAll() ([]*domain.Player, error) {
	panic("unimplemented")
}

// FindByID implements port.PlayerRepo.
func (*GameRepository) FindByID(id int) (*domain.Player, error) {
	panic("unimplemented")
}

// Update implements port.PlayerRepo.
func (*GameRepository) Update(entity domain.Player) error {
	panic("unimplemented")
}

func NewMongoGameRepository() port.PlayerRepo {
	return &GameRepository{}
}
