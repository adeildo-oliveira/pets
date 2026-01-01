package entities

import (
	"time"

	"pets/internal/models"
	entitiesdb "pets/internal/repository/entities_db"
	"pets/internal/utils"

)

type Pet struct {
	ID         int64
	Name       string
	Age        int
	Breed      string
	NickName   string
	Status     bool
	DateCreate time.Time
	DateUpdate time.Time
}

func NewPet(id int64, name string, age int, breed string, nickName string, status bool) *Pet {
	return &Pet{
		ID:         id,
		Name:       name,
		Age:        age,
		Breed:      breed,
		NickName:   utils.NullStringToString(utils.NullStringToSql(nickName)),
		Status:     status,
		DateCreate: time.Now(),
		DateUpdate: time.Now(),
	}
}

func MapToPetDB(pet *Pet) *entitiesdb.PetDB {
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

func MapModelToPetEntity(createPet *models.CreatePet) *Pet {
	return NewPet(0, createPet.Name, createPet.Age, createPet.Breed, createPet.NickName, true)
}

func MapToListEntityPet(petsDb []entitiesdb.PetDB) []Pet {
	pets := make([]Pet, len(petsDb))
	for i, petDb := range petsDb {
		pets[i] = *NewPet(petDb.ID, petDb.Name, petDb.Age, petDb.Breed, utils.NullStringToString(petDb.NickName), petDb.Status)
	}
	return pets
}

func (p *Pet) MapToPetResponse() *models.PetResponse {
	return &models.PetResponse{
		ID:       p.ID,
		Name:     p.Name,
		Age:      p.Age,
		Breed:    p.Breed,
		NickName: p.NickName,
		Status:   p.Status,
	}
}

func MapToListPetResponse(pets []Pet) []models.PetResponse {
	petsResponse := make([]models.PetResponse, len(pets))
	for i, pet := range pets {
		petsResponse[i] = *pet.MapToPetResponse()
	}
	return petsResponse
}
