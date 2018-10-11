package models

import "github.com/voonik/goFramework/pkg/database"

// ProductsGroupProductCollection ...
type ProductsGroupProductCollection struct {
	database.VModel
	ProductID                uint64
	GroupProductCollectionID uint64
}
