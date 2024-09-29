package repository

import (
	"calculate_product_packs/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMemoryPackSizeRepository(t *testing.T) {
	t.Run("Create repository with non-empty pack sizes", func(t *testing.T) {
		packSizes := []domain.PackSize{250, 500, 1000}
		repo := NewMemoryPackSizeRepository(packSizes)

		assert.NotNil(t, repo)
	})
}

func TestMemoryPackSizeRepository_GetPackSizes(t *testing.T) {
	t.Run("Repository with non-empty pack sizes", func(t *testing.T) {
		packSizes := []domain.PackSize{250, 500, 1000}
		repo := NewMemoryPackSizeRepository(packSizes)

		retrievedSizes := repo.GetPackSizes()

		assert.Equal(t, packSizes, retrievedSizes)
	})

	t.Run("Repository with empty pack sizes", func(t *testing.T) {
		emptyRepo := NewMemoryPackSizeRepository([]domain.PackSize{})

		emptySizes := emptyRepo.GetPackSizes()

		assert.NotNil(t, emptySizes)
	})
}
