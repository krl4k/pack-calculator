package http

import (
	"calculate_product_packs/internal/domain"
	"calculate_product_packs/internal/interfaces/http/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPackCalculatorHandler_CalculatePacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCalculator := mocks.NewMockPackCalculator(ctrl)

	handler := NewPackCalculatorHandler(mockCalculator)

	tests := []struct {
		name           string
		orderSize      string
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Valid order size",
			orderSize: "500",
			mockSetup: func() {
				mockCalculator.EXPECT().
					Execute(500).
					Return([]domain.PackResult{{Size: 500, Count: 1}}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `[{"Size":500,"Count":1}]
`,
		},
		{
			name:           "Invalid order size",
			orderSize:      "invalid",
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid order size\n",
		},
		{
			name:      "Order size must be greater than zero",
			orderSize: "0",
			mockSetup: func() {
				mockCalculator.EXPECT().
					Execute(0).
					Return(nil, domain.OrderSizeMustBeGreaterThanZeroError)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Order size must be greater than 0\n",
		},
		{
			name:      "No pack sizes available",
			orderSize: "100",
			mockSetup: func() {
				mockCalculator.EXPECT().
					Execute(100).
					Return(nil, domain.PackSizesNotFoundError)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "No pack sizes available\n",
		},
		{
			name:      "Internal server error",
			orderSize: "1000",
			mockSetup: func() {
				mockCalculator.EXPECT().
					Execute(1000).
					Return(nil, errors.New("unexpected error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "unexpected error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req, err := http.NewRequest("GET", "/calculate-packs?orderSize="+tt.orderSize, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.CalculatePacks(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestPackCalculatorHandler_CalculatePacks_JSONResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCalculator := mocks.NewMockPackCalculator(ctrl)
	handler := NewPackCalculatorHandler(mockCalculator)

	expectedResult := []domain.PackResult{
		{Size: 500, Count: 1},
		{Size: 250, Count: 1},
	}

	mockCalculator.EXPECT().
		Execute(750).
		Return(expectedResult, nil)

	req, err := http.NewRequest("GET", "/calculate-packs?orderSize=750", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.CalculatePacks(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var result []domain.PackResult
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
