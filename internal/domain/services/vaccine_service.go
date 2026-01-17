package services

import (
	"context"

	"pets/internal/domain/entities"
	"pets/internal/domain/interfaces"
)

type vaccineService struct {
	vaccineRepo interfaces.IVaccineRepository
}

func NewVaccineService(vaccineRepo interfaces.IVaccineRepository) interfaces.IVaccineService {
	return &vaccineService{vaccineRepo: vaccineRepo}
}

func (s *vaccineService) ListVaccines(ctx context.Context) ([]entities.Vaccine, error) {
	vaccinesDb, err := s.vaccineRepo.ListVaccines(ctx)
	if err != nil {
		return nil, err
	}
	vaccines := entities.MapToListEntityVaccine(vaccinesDb)
	return vaccines, nil
}

func (s *vaccineService) CreateVaccine(ctx context.Context, vaccine *entities.Vaccine) error {
	vaccineDb := entities.MapToVaccineDB(vaccine)

	err := s.vaccineRepo.CreateVaccine(ctx, vaccineDb)

	if err != nil {
		return err
	}
	return nil
}
