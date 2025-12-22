package entities

import (
	"time"

	entitiesdb "pets/internal/repository/entities_db"
	"pets/internal/utils"
)

type Pet struct {
	ID         string
	Name       string
	Age        int
	Breed      string
	NickName   string
	Status     bool
	DateCreate time.Time
	DateUpdate time.Time
}

func NewPet(entitiesdb *entitiesdb.PetDB) *Pet {
	return &Pet{
		ID:         entitiesdb.ID,
		Name:       entitiesdb.Name,
		Age:        entitiesdb.Age,
		Breed:      entitiesdb.Breed,
		NickName:   utils.NullStringToString(entitiesdb.NickName),
		Status:     entitiesdb.Status,
		DateCreate: utils.NullTimeToTime(entitiesdb.DateCreate),
		DateUpdate: utils.NullTimeToTime(entitiesdb.DateUpdate),
	}
}

func ToPets(entitiesdbs *[]entitiesdb.PetDB) *[]Pet {
	pets := make([]Pet, len(*entitiesdbs))

	for i, entitiesdb := range *entitiesdbs {
		pets[i] = *NewPet(&entitiesdb)
	}

	return &pets
}

func ToPetDB(pet *Pet) *entitiesdb.PetDB {
	return &entitiesdb.PetDB{
		ID:         pet.ID,
		Name:       pet.Name,
		Age:        pet.Age,
		Breed:      pet.Breed,
		NickName:   utils.NullStringToSql(pet.NickName),
		Status:     pet.Status,
		DateCreate: utils.NullTimeToSql(pet.DateCreate),
		DateUpdate: utils.NullTimeToSql(pet.DateUpdate),
	}
}
