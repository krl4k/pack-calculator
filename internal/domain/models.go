package domain

// PackSize represents the size of a pack as an integer
type PackSize int

// PackResult represents the result of a pack calculation,
// containing the size of the pack and the count needed
type PackResult struct {
	Size  PackSize
	Count int
}

// PackSizeRepository defines the interface for pack size storage operations.
// Using an interface here supports clean architecture principles and facilitates
// easier testing by allowing mock implementations in unit tests.
//
//go:generate mockgen -destination=mocks/mock_pack_size_repository.go -package=mocks calculate_product_packs/internal/domain PackSizeRepository
type PackSizeRepository interface {
	GetPackSizes() []PackSize
	UpdatePackSizes(sizes []PackSize) error
}
