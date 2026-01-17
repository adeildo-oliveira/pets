package models

import (
	"time"
)

type CreatePet struct {
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=15"`
	Breed    string `json:"breed" binding:"required"`
	NickName string `json:"nickName" binding:"required"`
}

type CreateVaccine struct {
	Name      string    `json:"nome" binding:"required"`
	DateGiven time.Time `json:"dataAplicacao" binding:"required,datetime=2006-01-02 15:00:00"`
	NextDue   time.Time `json:"proximaDose" binding:"required,datetime=2006-01-02 15:00:00"`
	IdPet     int64     `json:"idPet" binding:"required"`
}

type CreateConsulta struct {
	Date         time.Time `json:"data" binding:"required,datetime=2006-01-02 15:00:00"`
	Description  string    `json:"descricao" binding:"required"`
	Veterinarian string    `json:"veterinario" binding:"required"`
	Crmv         string    `json:"crmv" binding:"required"`
}

type CreateTutor struct {
	Name  string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"telefone" binding:"required"`
}

type CreateEndereco struct {
	Street     string `json:"rua" binding:"required"`
	Number     string `json:"numero" binding:"required"`
	City       string `json:"cidade" binding:"required"`
	State      string `json:"estado" binding:"required"`
	PostalCode string `json:"cep" binding:"required"`
}
