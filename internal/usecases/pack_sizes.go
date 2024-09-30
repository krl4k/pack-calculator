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
	if len(sizes) == 0 {
		return domain.EmptyPackSizesError
	}

	for _, size := range sizes {
		if size <= 0 {
			return domain.InvalidPackSizeError
		}
	}

	return uc.repo.UpdatePackSizes(sizes)
}

func (uc *PackSizesUseCase) GetPackSizes() []domain.PackSize {
	return uc.repo.GetPackSizes()
}
