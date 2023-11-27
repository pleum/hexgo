package mongo

import (
	"github.com/pleum/hexgo/internal/core/domain"
	"github.com/pleum/hexgo/internal/core/port"
)

type PlayerRepository struct {
}

// Create implements port.PlayerRepo.
func (PlayerRepository) Create(entity domain.Player) (int, error) {
	panic("unimplemented")
}

// DeleteByID implements port.PlayerRepo.
func (PlayerRepository) DeleteByID(id int) error {
	panic("unimplemented")
}

// FindAll implements port.PlayerRepo.
func (PlayerRepository) FindAll() ([]*domain.Player, error) {
	panic("unimplemented")
}

// FindByID implements port.PlayerRepo.
func (PlayerRepository) FindByID(id int) (*domain.Player, error) {
	panic("unimplemented")
}

// Update implements port.PlayerRepo.
func (PlayerRepository) Update(entity domain.Player) error {
	panic("unimplemented")
}

func NewMongoPlayerRepositry() port.PlayerRepo {
	return &PlayerRepository{}
}
