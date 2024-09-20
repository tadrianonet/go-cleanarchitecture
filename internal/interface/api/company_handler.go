package api

import (
	"api_company/internal/usecase"
	"context"
	"encoding/json"
	"net/http"
)

type CompanyHandler struct {
	Usecase usecase.CompanyUsecase
}

func NewCompanyHandler(u usecase.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{Usecase: u}
}

// GetAll godoc
// @Summary Retrieve all companies
// @Description Get all companies from the database
// @Tags company
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Company
// @Router /company [get]
func (h *CompanyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	companies, err := h.Usecase.GetAllCompanies(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(companies)
}
