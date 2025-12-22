package models

import "pets/internal/domain/entities"

type PetResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Breed    string `json:"breed"`
	NickName string `json:"nickName"`
	Status   bool   `json:"status"`
}

func NewPetResponse(entities entities.Pet) *PetResponse {
	return &PetResponse{
		ID:       entities.ID,
		Name:     entities.Name,
		Age:      entities.Age,
		Breed:    entities.Breed,
		NickName: entities.NickName,
		Status:   entities.Status,
	}
}

func Map[TRequest, TResponse any](items *[]TRequest, itemsToResponse func(TRequest) TResponse) *[]TResponse {
	listItems := make([]TResponse, len(*items))

	for i, item := range *items {
		listItems[i] = itemsToResponse(item)
	}

	return &listItems
}
