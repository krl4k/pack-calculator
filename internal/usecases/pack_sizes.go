package usecases

import (
	"calculate_product_packs/internal/domain"
)

type PackSizesUseCase struct {
	repo domain.PackSizeRepository
}

func NewPackSizesUseCase(repo domain.PackSizeRepository) *PackSizesUseCase {
	return &PackSizesUseCase{repo: repo}
}

func (uc *PackSizesUseCase) UpdatePackSizes(sizes []domain.PackSize) error {
	// Check for empty input
	if len(sizes) == 0 {
		return domain.EmptyPackSizesError
	}

	// Validate each pack size
	for _, size := range sizes {
		if size <= 0 {
			return domain.InvalidPackSizeError
		}
	}

	// Update pack sizes in repository
	return uc.repo.UpdatePackSizes(sizes)
}

func (uc *PackSizesUseCase) GetPackSizes() []domain.PackSize {
	// Retrieve pack sizes from repository
	return uc.repo.GetPackSizes()
}
