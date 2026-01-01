package interfaces

import (
	"context"

	"pets/internal/domain/entities"
)

type IPetService interface {
	ListPets(ctx context.Context) ([]entities.Pet, error)
	CreatePet(ctx context.Context, pet *entities.Pet) error
	UpdatePet(ctx context.Context, id int64, pet *entities.Pet) error
	DeletePet(ctx context.Context, id int64) error
}
