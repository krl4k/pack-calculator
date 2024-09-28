package usecases

import (
	"calculate_product_packs/internal/domain"
	"fmt"
	"sort"
)

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
		return nil, fmt.Errorf("order size must be greater than 0")
	}
	packSizes := uc.repo.GetPackSizes()
	if len(packSizes) == 0 {
		return nil, fmt.Errorf("no pack sizes available")
	}

	// Sort pack sizes in ascending order
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] < packSizes[j]
	})

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

func (uc *CalculatePacksUseCase) calculatePacks(orderSize int, packSizes []domain.PackSize) map[domain.PackSize]int {
	result := make(map[domain.PackSize]int)
	remaining := orderSize

	// Find the smallest pack size that can fulfill the order
	var smallestSufficientPack domain.PackSize
	for _, size := range packSizes {
		if size >= domain.PackSize(orderSize) {
			smallestSufficientPack = size
			break
		}
	}

	if smallestSufficientPack > 0 {
		// Check if using the smallest sufficient pack results in less waste
		wasteWithSmallestSufficient := int(smallestSufficientPack) - orderSize
		wasteWithSmallerPacks := uc.calculateWasteWithSmallerPacks(orderSize, packSizes)

		if wasteWithSmallestSufficient <= wasteWithSmallerPacks {
			result[smallestSufficientPack] = 1
			return result
		}
	}

	// Use combination of smaller packs
	for i := len(packSizes) - 1; i >= 0 && remaining > 0; i-- {
		if remaining >= int(packSizes[i]) {
			count := remaining / int(packSizes[i])
			result[packSizes[i]] = count
			remaining -= count * int(packSizes[i])
		}
	}

	// If there are still remaining items, add one more of the smallest pack that can cover it
	if remaining > 0 {
		for _, size := range packSizes {
			if size >= domain.PackSize(remaining) {
				result[size]++
				break
			}
		}
	}

	return result
}

func (uc *CalculatePacksUseCase) calculateWasteWithSmallerPacks(orderSize int, packSizes []domain.PackSize) int {
	remaining := orderSize
	for i := len(packSizes) - 1; i >= 0 && remaining > 0; i-- {
		count := remaining / int(packSizes[i])
		remaining -= count * int(packSizes[i])
	}
	if remaining > 0 {
		remaining = int(packSizes[0]) - remaining
	}
	return remaining
}
