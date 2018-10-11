package models

// GroupPrice ...
type GroupPrice struct {
	ProductID  uint64
	VariantID  uint64
	GroupPrice float64
}

// GroupPriceInterface ...
type GroupPriceInterface interface {
}
