package interfaces

import (
	"context"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"

	entitiesdb "pets/internal/repository/entities_db"

)

type IPetRepository interface {
	ListPets(ctx context.Context) ([]entitiesdb.PetDB, error)
	CreatePet(ctx context.Context, pet *entitiesdb.PetDB) error
	UpdatePet(ctx context.Context, pet *entitiesdb.PetDB) error
}

type ISqlDB interface {
	GetSqlConnection(ctx context.Context) (*sqlx.DB, error)
}
