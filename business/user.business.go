package business

import (
	"main/access"
	"main/models"
)

type User struct {
	UserAccess *access.User
}

func (b *User) getUserById(id string) (*models.User, error) {
	return b.UserAccess.GetUserById(id)
}
