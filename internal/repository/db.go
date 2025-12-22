package repository

import (
	"context"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"pets/internal/config"
	"pets/internal/domain/interfaces"
)

type petDb struct {
	cfgDb  *config.Server
	logger *zap.Logger
}

func NewPetDb(cfgDb *config.Server, logger *zap.Logger) interfaces.ISqlDB {
	return &petDb{
		cfgDb:  cfgDb,
		logger: logger,
	}
}

func (r *petDb) GetSqlConnection(ctx context.Context) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		r.cfgDb.DataBase.User,
		r.cfgDb.DataBase.Password,
		r.cfgDb.DataBase.Host,
		r.cfgDb.DataBase.Port,
		r.cfgDb.DataBase.Name,
	)

	db, err := sqlx.Connect("sqlserver", dsn)
	if err != nil {
		r.logger.Error("Error while connecting to the database", zap.Error(err))
		return nil, err
	}

	return db, nil
}
