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

func (s *petService) ListPets(ctx context.Context) ([]entities.Pet, error) {
	petsDb, err := s.repo.ListPets(ctx)

	if err != nil {
		return nil, err
	}

	return entities.MapToListEntityPet(petsDb), nil
}

func (s *petService) CreatePet(ctx context.Context, pet *entities.Pet) error {
	petDb := entities.MapToPetDB(pet)
	err := s.repo.CreatePet(ctx, petDb)

	if err != nil {
		return err
	}

	return nil
}

func (s *petService) UpdatePet(ctx context.Context, id int64, pet *entities.Pet) error {
	petDb := entities.MapToPetDB(pet)
	petDb.ID = id
	err := s.repo.UpdatePet(ctx, petDb)

	if err != nil {
		return err
	}

	return nil
}

func (s *petService) DeletePet(ctx context.Context, id int64) error {
	err := s.repo.DeletePet(ctx, id)
	if err != nil {
		return err
	}

	return nil
}