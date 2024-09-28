package usecases

import (
	"calculate_product_packs/internal/domain"
	"fmt"
	"sort"
)

//go:generate mockgen -destination=./mocks/mock_pack_size_repository.go -package=mocks calculate_product_packs/internal/usecases PackSizeRepository
type PackSizeRepository interface {
	GetPackSizes() []domain.PackSize
}

type CalculatePacksUseCase struct {
	repo PackSizeRepository
}

func NewCalculatePacksUseCase(repo PackSizeRepository) *CalculatePacksUseCase {
	return &CalculatePacksUseCase{repo: repo}
}

func (uc *CalculatePacksUseCase) Execute(orderSize int) ([]domain.PackResult, error) {
	if orderSize <= 0 {
		return []domain.PackResult{}, fmt.Errorf("order size must be greater than 0")
	}
	packSizes := uc.repo.GetPackSizes()
	if len(packSizes) == 0 {
		return []domain.PackResult{}, fmt.Errorf("no pack sizes available")
	}

	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})

	result := uc.calculatePacks(orderSize, packSizes)

	var packResults []domain.PackResult
	for size, count := range result {
		packResults = append(packResults, domain.PackResult{Size: size, Count: count})
	}

	sort.Slice(packResults, func(i, j int) bool {
		return packResults[i].Size > packResults[j].Size
	})
	return packResults, nil
}

func (uc *CalculatePacksUseCase) calculatePacks(orderSize int, packSizes []domain.PackSize) map[domain.PackSize]int {
	result := make(map[domain.PackSize]int)
	remaining := orderSize

	for _, size := range packSizes {
		if remaining >= int(size) {
			count := remaining / int(size)
			result[size] = count
			remaining -= count * int(size)
		}
	}

	// If there are remaining items, add one more of the smallest pack
	if remaining > 0 {
		smallestPack := packSizes[len(packSizes)-1]
		result[smallestPack]++
	}

	return result
}
