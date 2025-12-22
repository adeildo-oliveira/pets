package repository

import (
	"context"

	"go.uber.org/zap"

	"pets/internal/domain/interfaces"
	entitiesdb "pets/internal/repository/entities_db"
	"pets/internal/utils"
)

type petRepository struct {
	db     interfaces.ISqlDB
	logger *zap.Logger
}

func NewPetRepository(db interfaces.ISqlDB, logger *zap.Logger) interfaces.IPetRepository {
	return &petRepository{db: db,
		logger: logger,
	}
}

func (r *petRepository) ListPets(ctx context.Context) (*[]entitiesdb.PetDB, error) {
	db, err := r.db.GetSqlConnection(ctx)

	if err != nil {
		r.logger.Error("Error while get connection pets", zap.Error(err))
		return nil, err
	}

	var pets []entitiesdb.PetDB

	err = db.Select(&pets, selectPetsQuery)
	if err != nil {
		r.logger.Error("Error while selecting pets", zap.Error(err))
		return nil, err
	}
	return &pets, nil
}

func (r *petRepository) CreatePet(ctx context.Context, pet *entitiesdb.PetDB) error {
	db, err := r.db.GetSqlConnection(ctx)
	if err != nil {
		r.logger.Error("Error while get connection pets", zap.Error(err))
		return err
	}

	_, err = db.NamedExec(insertPetQuery, pet)
	if err != nil {
		if utils.IsUniqueViolation(err) {
			r.logger.Error("Unique constraint violation", zap.Error(err))

			return err
		}

		r.logger.Error("Error while inserting pet", zap.Error(err))
		return err
	}
	return nil
}

func (r *petRepository) UpdatePet(ctx context.Context, pet *entitiesdb.PetDB) error {
	db, err := r.db.GetSqlConnection(ctx)
	if err != nil {
		r.logger.Error("Error while get connection pets", zap.Error(err))
		return err
	}

	_, err = db.NamedExec(updatePetQuery, pet)
	if err != nil {
		if utils.IsUniqueViolation(err) {
			r.logger.Error("Unique constraint violation", zap.Error(err))

			return err
		}

		r.logger.Error("Error while updating pet", zap.Error(err))
		return err
	}
	return nil
}
