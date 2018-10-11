package handlers

import (
	"github.com/voonik/gbs/internal/app/services"
)

// GetGameGroupInstance return the GameGroupService struct
func GetGameGroupInstance() *services.GameGroupService {
	return &services.GameGroupService{}
}
