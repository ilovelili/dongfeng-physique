package controllers

import (
	"github.com/ilovelili/dongfeng-physique/services/server/core/models"
	"github.com/ilovelili/dongfeng-physique/services/server/core/repositories"
)

// UserController user profile controller
type UserController struct {
	repository *repositories.UserRepository
}

// NewUserController new controller
func NewUserController() *UserController {
	return &UserController{
		repository: repositories.NewUserRepository(),
	}
}

// GetUserByEmail get user by mail
func (c *UserController) GetUserByEmail(email string) (*models.User, error) {
	return c.repository.SelectByMail(email)
}
