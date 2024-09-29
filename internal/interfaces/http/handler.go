package http

import (
	"calculate_product_packs/internal/domain"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

//go:generate mockgen -destination=mocks/mock_pack_calculator.go -package=mocks calculate_product_packs/internal/interfaces/http PackCalculator
type PackCalculator interface {
	Execute(orderSize int) ([]domain.PackResult, error)
}

type PackCalculatorHandler struct {
	packCalculator PackCalculator
}

func NewPackCalculatorHandler(packCalculator PackCalculator) *PackCalculatorHandler {
	return &PackCalculatorHandler{packCalculator: packCalculator}
}

func (h *PackCalculatorHandler) CalculatePacks(w http.ResponseWriter, r *http.Request) {
	orderSize, err := strconv.Atoi(r.URL.Query().Get("orderSize"))
	if err != nil {
		http.Error(w, "Invalid order size", http.StatusBadRequest)
		return
	}

	result, err := h.packCalculator.Execute(orderSize)
	if err != nil {
		switch {
		case errors.Is(err, domain.OrderSizeMustBeGreaterThanZeroError):
			http.Error(w, "Order size must be greater than 0", http.StatusBadRequest)
		case errors.Is(err, domain.PackSizesNotFoundError):
			http.Error(w, "No pack sizes available", http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
