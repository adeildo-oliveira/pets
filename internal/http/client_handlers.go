package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"pets/internal/domain"
	"pets/internal/domain/entities"
	"pets/internal/domain/interfaces"
	"pets/internal/models"
)

type ClientHandler struct {
	petService interfaces.IPetService
}

func NewClientHandler(petService interfaces.IPetService) *ClientHandler {
	return &ClientHandler{
		petService: petService,
	}
}
func (h *ClientHandler) List(w http.ResponseWriter, r *http.Request) {
	petsService, err := h.petService.ListPets(r.Context())

	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusInternalServerError, "erro ao listar pets")
		return
	}

	response := domain.Response{
		Data: entities.MapToListPetResponse(petsService),
	}

	response.WriteJSON(w, http.StatusOK)
}

func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var petRequest models.CreatePet

	err := json.NewDecoder(r.Body).Decode(&petRequest)
	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	petEntity := entities.MapModelToPetEntity(&petRequest)
	err = h.petService.CreatePet(r.Context(), petEntity)
	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusInternalServerError, "erro ao criar pet")
		return
	}

	response := domain.Response{
		Data: "pet created successfully",
	}

	response.WriteJSON(w, http.StatusCreated)
}

func (h *ClientHandler) Update(w http.ResponseWriter, r *http.Request) {
	var petRequest models.CreatePet
	petId := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&petRequest)

	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	petEntity := entities.MapModelToPetEntity(&petRequest)

	id, err := strconv.ParseInt(petId, 10, 64)
	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	err = h.petService.UpdatePet(r.Context(), id, petEntity)

	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusInternalServerError, "erro ao atualizar pet")
		return
	}

	response := domain.Response{
		Data: "pet updated successfully",
	}
	response.WriteJSON(w, http.StatusOK)
}

func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(petId, 10, 64)
	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	err = h.petService.DeletePet(r.Context(), id)
	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusInternalServerError, "erro ao deletar pet")
		return
	}

	response := domain.Response{
		Data: "pet deleted successfully",
	}
	response.WriteJSON(w, http.StatusOK)
}
