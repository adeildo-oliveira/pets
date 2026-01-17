package http

import (
	"encoding/json"
	"net/http"

	"pets/internal/domain"
	"pets/internal/domain/entities"
	"pets/internal/domain/interfaces"
	"pets/internal/models"

)

type VaccineHandler struct {
	vaccineService interfaces.IVaccineService
}

func NewVaccineHandler(vaccineService interfaces.IVaccineService) *VaccineHandler {
	return &VaccineHandler{
		vaccineService: vaccineService,
	}
}

func (h *VaccineHandler) List(w http.ResponseWriter, r *http.Request) {
	vaccinesService, err := h.vaccineService.ListVaccines(r.Context())
	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusInternalServerError, "erro ao listar vacinas")
		return
	}
	response := domain.Response{
		Data: entities.MapToListVaccineResponse(vaccinesService),
	}
	response.WriteJSON(w, http.StatusOK)
}

func (h *VaccineHandler) Create(w http.ResponseWriter, r *http.Request) {
	var vaccineRequest models.CreateVaccine
	err := json.NewDecoder(r.Body).Decode(&vaccineRequest)
	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	vaccineEntity := entities.MapModelToVaccineEntity(&vaccineRequest)
	err = h.vaccineService.CreateVaccine(r.Context(), vaccineEntity)
	if err != nil {
		response := domain.Response{}
		response.WriteError(w, http.StatusInternalServerError, "erro ao criar vacina")
		return
	}
	response := domain.Response{
		Data: "vacina criada com sucesso",
	}
	response.WriteJSON(w, http.StatusCreated)
}
