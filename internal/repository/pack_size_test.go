package repository

import (
	"calculate_product_packs/internal/domain"
	"reflect"
	"testing"
)

func TestNewMemoryPackSizeRepository(t *testing.T) {
	t.Run("Create repository with non-empty pack sizes", func(t *testing.T) {
		packSizes := []domain.PackSize{250, 500, 1000}
		repo := NewMemoryPackSizeRepository(packSizes)

		if repo == nil {
			t.Error("Expected non-nil repository, got nil")
		}
	})

	t.Run("Create repository with empty pack sizes", func(t *testing.T) {
		emptyRepo := NewMemoryPackSizeRepository([]domain.PackSize{})

		if emptyRepo == nil {
			t.Error("Expected non-nil repository even with empty pack sizes, got nil")
		}
	})
}

func TestMemoryPackSizeRepository_GetPackSizes(t *testing.T) {
	t.Run("Repository with non-empty pack sizes", func(t *testing.T) {
		packSizes := []domain.PackSize{250, 500, 1000}
		repo := NewMemoryPackSizeRepository(packSizes)

		retrievedSizes := repo.GetPackSizes()

		if !reflect.DeepEqual(retrievedSizes, packSizes) {
			t.Errorf("Expected pack sizes %v, got %v", packSizes, retrievedSizes)
		}
	})

	t.Run("Repository with empty pack sizes", func(t *testing.T) {
		emptyRepo := NewMemoryPackSizeRepository([]domain.PackSize{})

		emptySizes := emptyRepo.GetPackSizes()

		if len(emptySizes) != 0 {
			t.Errorf("Expected empty pack sizes, got %v", emptySizes)
		}
	})
}
