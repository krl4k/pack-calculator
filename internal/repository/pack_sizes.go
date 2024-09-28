package repository

import (
	"calculate_product_packs/internal/domain"
	"calculate_product_packs/internal/usecases"
)

type MemoryPackSizeRepository struct {
	packSizes []domain.PackSize
}

func NewMemoryPackSizeRepository(packSizes []domain.PackSize) usecases.PackSizeRepository {
	return &MemoryPackSizeRepository{
		packSizes: packSizes,
	}
}

func (r *MemoryPackSizeRepository) GetPackSizes() []domain.PackSize {
	return r.packSizes
}
