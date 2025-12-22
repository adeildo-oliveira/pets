package entitiesdb

import (
	"database/sql"
)

type PetDB struct {
	ID         string         `db:"id"`
	Name       string         `db:"name_pet"`
	Age        int            `db:"age_pet"`
	Breed      string         `db:"breed_pet"`
	NickName   sql.NullString `db:"nick_name_pet"`
	Status     bool           `db:"status_pet"`
	DateCreate sql.NullTime   `db:"date_create_pet"`
	DateUpdate sql.NullTime   `db:"date_update_pet"`
}
