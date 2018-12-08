package repositories

import (
	"github.com/ilovelili/dongfeng-physique/services/server/core/models"
)

// UserRepository user info repository
type UserRepository struct{}

// NewUserRepository init UserProfile repository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// SelectByMail select user by mail
func (r *UserRepository) SelectByMail(email string) (user *models.User, err error) {
	var u models.User
	query := Table("users").Alias("u").Where().Eq("u.email", email).Eq("u.enabled", 1).Sql()
	if err = session().Find(query, nil).Single(&u); err != nil {
		return
	}
	user = &u
	return
}
