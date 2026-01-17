package repository

import (
	"context"

	"go.uber.org/zap"

	"pets/internal/domain/interfaces"
	entitiesdb "pets/internal/repository/entities_db"
)

type vaccineRepository struct {
	db     interfaces.ISqlDB
	logger *zap.Logger
}

func NewVaccineRepository(db interfaces.ISqlDB, logger *zap.Logger) interfaces.IVaccineRepository {
	return &vaccineRepository{db: db,
		logger: logger,
	}
}

func (r *vaccineRepository) ListVaccines(ctx context.Context) ([]entitiesdb.VaccineDB, error) {
	db, err := r.db.GetSqlConnection(ctx)

	if err != nil {
		r.logger.Error("Error while get connection vaccines", zap.Error(err))
		return nil, err
	}

	var vaccines []entitiesdb.VaccineDB
	err = db.Select(&vaccines, selectVaccinesQuery)

	if err != nil {
		r.logger.Error("Error while selecting vaccines", zap.Error(err))
		return nil, err
	}

	return vaccines, nil
}

func (r *vaccineRepository) CreateVaccine(ctx context.Context, vaccine *entitiesdb.VaccineDB) error {
	db, err := r.db.GetSqlConnection(ctx)

	if err != nil {
		r.logger.Error("Error while get connection vaccines", zap.Error(err))
		return err
	}
	_, err = db.NamedExec(insertVaccineQuery, vaccine)

	if err != nil {
		r.logger.Error("Error while inserting vaccine", zap.Error(err))
		return err
	}

	return nil
}
