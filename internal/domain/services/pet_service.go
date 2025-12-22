package services

import (
	"context"

	"pets/internal/domain/entities"
	"pets/internal/domain/interfaces"
)

type petService struct {
	repo interfaces.IPetRepository
}

func NewPetService(repo interfaces.IPetRepository) interfaces.IPetService {
	return &petService{repo: repo}
}

func (s *petService) ListPets(ctx context.Context) (*[]entities.Pet, error) {
	petsDb, err := s.repo.ListPets(ctx)

	if err != nil {
		return nil, err
	}

	pets := entities.ToPets(petsDb)
	return pets, nil
}

func (s *petService) CreatePet(ctx context.Context, pet *entities.Pet) error {
	petDb := entities.ToPetDB(pet)
	err := s.repo.CreatePet(ctx, petDb)

	if err != nil {
		return err
	}

	return nil
}

func (s *petService) UpdatePet(ctx context.Context, id string, pet *entities.Pet) error {
	petDb := entities.ToPetDB(pet)
	petDb.ID = id
	err := s.repo.UpdatePet(ctx, petDb)

	if err != nil {
		return err
	}

	return nil
}
