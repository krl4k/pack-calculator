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

//go:generate mockgen -destination=mocks/mock_pack_sizer.go -package=mocks calculate_product_packs/internal/interfaces/http PackSizer
type PackSizer interface {
	UpdatePackSizes(sizes []domain.PackSize) error
	GetPackSizes() []domain.PackSize
}

type PackCalculatorHandler struct {
	packCalculator   PackCalculator
	packSizesUseCase PackSizer
}

func NewPackCalculatorHandler(
	packCalculator PackCalculator,
	packSizesUseCase PackSizer) *PackCalculatorHandler {
	return &PackCalculatorHandler{
		packCalculator:   packCalculator,
		packSizesUseCase: packSizesUseCase,
	}
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

func (h *PackCalculatorHandler) UpdatePackSizes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var sizes []domain.PackSize
	err := json.NewDecoder(r.Body).Decode(&sizes)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.packSizesUseCase.UpdatePackSizes(sizes)
	if err != nil {
		switch err {
		case domain.EmptyPackSizesError, domain.InvalidPackSizeError:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "Failed to update pack sizes", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pack sizes updated successfully"))
}

func (h *PackCalculatorHandler) GetPackSizes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sizes := h.packSizesUseCase.GetPackSizes()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sizes)
}
