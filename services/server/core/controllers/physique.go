package controllers

import (
	"github.com/ilovelili/dongfeng-physique/services/server/core/repositories"
)

// PhysiqueController notification controller
type PhysiqueController struct {
	repository *repositories.NotificationRepository
}

// NewPhysiqueController new controller
func NewPhysiqueController() *PhysiqueController {
	return &PhysiqueController{
		repository: repositories.NewNotificationRepository(),
	}
}
