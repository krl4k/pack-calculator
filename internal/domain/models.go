package domain

type PackSize int

type PackResult struct {
	Size  PackSize
	Count int
}

type PackSizeRepository interface {
	GetPackSizes() []PackSize
	UpdatePackSizes(sizes []PackSize) error
}
