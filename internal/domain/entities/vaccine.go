package entities

import (
	"time"

	"pets/internal/models"
	entitiesdb "pets/internal/repository/entities_db"
	"pets/internal/utils"
)

type Vaccine struct {
	ID         int64
	Name       string
	DateGiven  time.Time
	NextDue    time.Time
	IdPet      int64
	Status     bool
	DateCreate time.Time
	DateUpdate time.Time
}

func NewVaccine(id int64, name string, dateGiven time.Time, nextDue time.Time, idPet int64, status bool) *Vaccine {
	return &Vaccine{
		ID:         id,
		Name:       name,
		DateGiven:  dateGiven,
		NextDue:    nextDue,
		IdPet:      idPet,
		Status:     status,
		DateCreate: time.Now(),
		DateUpdate: time.Now(),
	}
}

func MapModelToVaccineEntity(createVaccine *models.CreateVaccine) *Vaccine {
	return NewVaccine(0, createVaccine.Name, createVaccine.DateGiven, createVaccine.NextDue, createVaccine.IdPet, true)
}

func MapToVaccineDB(vaccine *Vaccine) *entitiesdb.VaccineDB {
	return &entitiesdb.VaccineDB{
		ID:         vaccine.ID,
		Name:       vaccine.Name,
		DateGiven:  utils.NullTimeToSql(vaccine.DateGiven),
		NextDue:    utils.NullTimeToSql(vaccine.NextDue),
		IdPet:      vaccine.IdPet,
		Status:     vaccine.Status,
		DateCreate: utils.NullTimeToSql(vaccine.DateCreate),
		DateUpdate: utils.NullTimeToSql(vaccine.DateUpdate),
	}
}

func MapToListEntityVaccine(vaccinesDb []entitiesdb.VaccineDB) []Vaccine {
	vaccines := make([]Vaccine, len(vaccinesDb))

	for i, vaccineDb := range vaccinesDb {
		vaccines[i] = *NewVaccine(vaccineDb.ID,
			vaccineDb.Name,
			utils.NullTimeToTime(vaccineDb.DateGiven),
			utils.NullTimeToTime(vaccineDb.NextDue),
			vaccineDb.IdPet,
			vaccineDb.Status)
	}

	return vaccines
}

func (v *Vaccine) MapToVaccineResponse() *models.VaccineResponse {
	return &models.VaccineResponse{
		ID:        v.ID,
		Name:      v.Name,
		DateGiven: v.DateGiven,
		NextDue:   v.NextDue,
	}
}

func MapToListVaccineResponse(vaccines []Vaccine) []models.VaccineResponse {
	vaccineResponses := make([]models.VaccineResponse, len(vaccines))
	for i, vaccine := range vaccines {
		vaccineResponses[i] = *vaccine.MapToVaccineResponse()
	}
	return vaccineResponses
}
