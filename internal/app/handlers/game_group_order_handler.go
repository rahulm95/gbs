package handlers

import "github.com/voonik/gbs/internal/app/services"

// GetGameGroupOrderInstance ...
func GetGameGroupOrderInstance() *services.GameGroupOrderService {
	return &services.GameGroupOrderService{}
}
