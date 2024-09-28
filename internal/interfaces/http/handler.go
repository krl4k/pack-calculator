package http

import (
	"calculate_product_packs/internal/usecases"
	"encoding/json"
	"net/http"
	"strconv"
)

type PackCalculatorHandler struct {
	useCase *usecases.CalculatePacksUseCase
}

func NewPackCalculatorHandler(useCase *usecases.CalculatePacksUseCase) *PackCalculatorHandler {
	return &PackCalculatorHandler{useCase: useCase}
}

func (h *PackCalculatorHandler) CalculatePacks(w http.ResponseWriter, r *http.Request) {
	orderSize, err := strconv.Atoi(r.URL.Query().Get("orderSize"))
	if err != nil {
		http.Error(w, "Invalid order size", http.StatusBadRequest)
		return
	}

	result, err := h.useCase.Execute(orderSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
