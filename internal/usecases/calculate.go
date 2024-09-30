package usecases

import (
	"calculate_product_packs/internal/domain"
	"sort"
)

type CalculatePacksUseCase struct {
	repo domain.PackSizeRepository
}

func NewCalculatePacksUseCase(repo domain.PackSizeRepository) *CalculatePacksUseCase {
	return &CalculatePacksUseCase{repo: repo}
}

func (uc *CalculatePacksUseCase) Execute(orderSize int) ([]domain.PackResult, error) {
	if orderSize <= 0 {
		return nil, domain.OrderSizeMustBeGreaterThanZeroError
	}

	packSizes := uc.repo.GetPackSizes()
	if len(packSizes) == 0 {
		return nil, domain.PackSizesNotFoundError
	}

	result := uc.calculatePacks(orderSize, packSizes)

	var packResults []domain.PackResult
	for size, count := range result {
		if count > 0 {
			packResults = append(packResults, domain.PackResult{Size: size, Count: count})
		}
	}

	// Sort results by pack size in descending order
	sort.Slice(packResults, func(i, j int) bool {
		return packResults[i].Size > packResults[j].Size
	})
	return packResults, nil
}

func (uc *CalculatePacksUseCase) calculatePacks(items int, packSizes []domain.PackSize) map[domain.PackSize]int {
	// Sort pack sizes in ascending order
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] < packSizes[j]
	})

	necessaryPacks := make(map[domain.PackSize]int)
	lastUsedPackIndex := len(packSizes) - 1
	diff := 0

	for lastUsedPackIndex > 0 {
		if items-int(packSizes[lastUsedPackIndex]) >= 0 {
			necessaryPacks[packSizes[lastUsedPackIndex]]++
			items -= int(packSizes[lastUsedPackIndex])
		} else {
			if _, exists := necessaryPacks[packSizes[lastUsedPackIndex]]; exists {
				diff = int(packSizes[lastUsedPackIndex]) - items
				if int(packSizes[lastUsedPackIndex-1]) > diff {
					necessaryPacks[packSizes[lastUsedPackIndex]]++
					items -= int(packSizes[lastUsedPackIndex])
					break
				}
			}
			lastUsedPackIndex--
		}
	}

	if items > 0 {
		for _, size := range packSizes {
			if int(size) >= items {
				necessaryPacks[size]++
				items -= int(size)
				break
			}
		}
	}

	for items > 0 {
		necessaryPacks[packSizes[lastUsedPackIndex]]++
		items -= int(packSizes[lastUsedPackIndex])
	}

	return necessaryPacks
}
