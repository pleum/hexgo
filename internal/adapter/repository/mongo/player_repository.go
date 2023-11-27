package mongo

import (
	"github.com/pleum/hexgo/internal/core/domain"
	"github.com/pleum/hexgo/internal/core/port"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerRepository struct {
	mongoClient *mongo.Client
}

// CreateMany implements port.PlayerRepo.
func (*PlayerRepository) CreateMany(entities []domain.Player) ([]domain.Player, error) {
	return []domain.Player{}, nil
}

// Create implements port.PlayerRepo.
func (PlayerRepository) Create(entity domain.Player) (domain.Player, error) {
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

func NewMongoPlayerRepositry(mongoClient *mongo.Client) port.PlayerRepo {
	return &PlayerRepository{
		mongoClient: mongoClient,
	}
}
