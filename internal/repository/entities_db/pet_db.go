package entitiesdb

import (
	"database/sql"

)

type PetDB struct {
	ID         int64          `db:"id"`
	Name       string         `db:"name_pet"`
	Age        int            `db:"age_pet"`
	Breed      string         `db:"breed_pet"`
	NickName   sql.NullString `db:"nick_name_pet"`
	Status     bool           `db:"status_pet"`
	DateCreate sql.NullTime   `db:"date_create_pet"`
	DateUpdate sql.NullTime   `db:"date_update_pet"`
}

type VaccineDB struct {
	ID         int64        `db:"id"`
	Name       string       `db:"name_vaccine"`
	DateGiven  sql.NullTime `db:"date_given_vaccine"`
	NextDue    sql.NullTime `db:"next_due_vaccine"`
	IdPet      int64        `db:"id_pet"`
	Status     bool         `db:"status_vaccine"`
	DateCreate sql.NullTime `db:"date_create_vaccine"`
	DateUpdate sql.NullTime `db:"date_update_vaccine"`
}
