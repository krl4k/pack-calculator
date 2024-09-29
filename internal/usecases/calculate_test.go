package usecases

import (
	"calculate_product_packs/internal/domain"
	"calculate_product_packs/internal/usecases/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePacksUseCase_Execute(t *testing.T) {
	tests := []struct {
		name          string
		packSizes     []domain.PackSize
		orderSize     int
		expectedPacks []domain.PackResult
		expectedError error
	}{
		{
			name:          "Simple case 1",
			packSizes:     []domain.PackSize{250, 500, 1000, 2000, 5000},
			orderSize:     1,
			expectedPacks: []domain.PackResult{{Size: 250, Count: 1}},
			expectedError: nil,
		},
		{
			name:          "Simple case 2",
			packSizes:     []domain.PackSize{250, 500, 1000, 2000, 5000},
			orderSize:     250,
			expectedPacks: []domain.PackResult{{Size: 250, Count: 1}},
			expectedError: nil,
		},
		{
			name:          "Simple case 3",
			packSizes:     []domain.PackSize{250, 500, 1000, 2000, 5000},
			orderSize:     251,
			expectedPacks: []domain.PackResult{{Size: 500, Count: 1}},
			expectedError: nil,
		},
		{
			name:          "Simple case 4",
			packSizes:     []domain.PackSize{250, 500, 1000, 2000, 5000},
			orderSize:     501,
			expectedPacks: []domain.PackResult{{Size: 500, Count: 1}, {Size: 250, Count: 1}},
			expectedError: nil,
		}, {
			name:          "Simple case 5",
			packSizes:     []domain.PackSize{250, 500, 1000, 2000, 5000},
			orderSize:     12001,
			expectedPacks: []domain.PackResult{{Size: 5000, Count: 2}, {Size: 2000, Count: 1}, {Size: 250, Count: 1}},
			expectedError: nil,
		},
		{
			name:          "Exact match",
			packSizes:     []domain.PackSize{250, 500, 1000, 2000, 5000},
			orderSize:     500,
			expectedPacks: []domain.PackResult{{Size: 500, Count: 1}},
			expectedError: nil,
		},
		{
			name:          "Multiple packs",
			packSizes:     []domain.PackSize{250, 500, 1000, 2000, 5000},
			orderSize:     12001,
			expectedPacks: []domain.PackResult{{Size: 5000, Count: 2}, {Size: 2000, Count: 1}, {Size: 250, Count: 1}},
			expectedError: nil,
		},
		{
			name:          "Large order size",
			packSizes:     []domain.PackSize{250, 500, 1000, 2000, 5000},
			orderSize:     1000000,
			expectedPacks: []domain.PackResult{{Size: 5000, Count: 200}},
			expectedError: nil,
		},
		{
			name:          "Single pack size",
			packSizes:     []domain.PackSize{1000},
			orderSize:     2500,
			expectedPacks: []domain.PackResult{{Size: 1000, Count: 3}},
			expectedError: nil,
		},
		{
			name:          "No pack sizes available",
			packSizes:     []domain.PackSize{},
			orderSize:     1000,
			expectedPacks: nil,
			expectedError: fmt.Errorf("no pack sizes available"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockPackSizeRepository(ctrl)
			mockRepo.EXPECT().GetPackSizes().Return(tt.packSizes)

			useCase := NewCalculatePacksUseCase(mockRepo)

			result, err := useCase.Execute(tt.orderSize)

			assert.Equal(t, tt.expectedError, err)

			assert.Equal(t, tt.expectedPacks, result)
		})
	}
}

func TestCalculatePacksUseCase_Execute_SortsPackSizes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPackSizeRepository(ctrl)
	mockRepo.EXPECT().GetPackSizes().Return([]domain.PackSize{250, 1000, 500, 5000, 2000})

	useCase := NewCalculatePacksUseCase(mockRepo)

	result, err := useCase.Execute(12001)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []domain.PackResult{{Size: 5000, Count: 2}, {Size: 2000, Count: 1}, {Size: 250, Count: 1}}
	assert.Equal(t, expected, result)
}

func TestCalculatePacksUseCase_Execute_EmptyPackSizes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPackSizeRepository(ctrl)
	mockRepo.EXPECT().GetPackSizes().Return([]domain.PackSize{})

	useCase := NewCalculatePacksUseCase(mockRepo)

	result, err := useCase.Execute(1000)

	assert.Error(t, err)
	assert.Equal(t, domain.PackSizesNotFoundError, err)
	assert.Empty(t, result)
}

func TestCalculatePacksUseCase_Execute_InvalidInputs(t *testing.T) {
	tests := []struct {
		name      string
		packSizes []domain.PackSize
		orderSize int
	}{
		{
			name:      "Negative order size",
			packSizes: []domain.PackSize{250, 500, 1000},
			orderSize: -100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockPackSizeRepository(ctrl)

			useCase := NewCalculatePacksUseCase(mockRepo)

			result, err := useCase.Execute(tt.orderSize)

			assert.Error(t, err)
			assert.Empty(t, result)
		})
	}
}
