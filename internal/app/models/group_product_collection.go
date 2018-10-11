package models

import "github.com/voonik/goFramework/pkg/database"

// GroupProductCollection ...
type GroupProductCollection struct {
	database.VModel
	IsActive bool
}

// GroupProductCollectionInterface ...
type GroupProductCollectionInterface interface {
}
