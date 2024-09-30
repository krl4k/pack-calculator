package repository

import (
	"calculate_product_packs/internal/domain"
	"sync"
)

type MemoryPackSizeRepository struct {
	lock      sync.RWMutex
	packSizes []domain.PackSize
}

func NewMemoryPackSizeRepository(packSizes []domain.PackSize) domain.PackSizeRepository {
	return &MemoryPackSizeRepository{
		lock:      sync.RWMutex{},
		packSizes: packSizes,
	}
}

func (r *MemoryPackSizeRepository) GetPackSizes() []domain.PackSize {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.packSizes
}

func (r *MemoryPackSizeRepository) UpdatePackSizes(sizes []domain.PackSize) error {
	r.lock.Lock()
	r.packSizes = sizes
	r.lock.Unlock()
	return nil
}
