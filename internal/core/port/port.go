package port

import "github.com/pleum/hexgo/internal/core/domain"

type Repo[T domain.Model] interface {
	Create(entity T) (int, error)
	Update(entity T) error
	FindByID(id int) (*T, error)
	FindAll() ([]*T, error)
	DeleteByID(id int) error
}

type GameRepo interface {
	Repo[domain.Game]
}

type PlayerRepo interface {
	Repo[domain.Player]
}
