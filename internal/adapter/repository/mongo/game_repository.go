package mongo

import (
	"github.com/pleum/hexgo/internal/core/domain"
	"github.com/pleum/hexgo/internal/core/port"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepository struct {
	mongoClient *mongo.Client
}

// CreateMany implements port.GameRepo.
func (*GameRepository) CreateMany(entities []domain.Game) ([]domain.Game, error) {
	panic("unimplemented")
}

// Create implements port.GameRepo.
func (*GameRepository) Create(entity domain.Game) (domain.Game, error) {
	return domain.Game{}, nil
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

func NewMongoGameRepository(mongoClient *mongo.Client) port.GameRepo {
	return &GameRepository{
		mongoClient: mongoClient,
	}
}
