package user

import (
	"github.com/shobky/giggs/config/database"
	"github.com/shobky/giggs/model"
)

func FindUserByEmail(email string) (*model.User, bool) {
	var user *model.User
	database.DB.First(&user, "email = ?", email)

	if user.ID == 0 {
		return nil, false
	}

	return user, true
}

func FindUserById(id string) (*model.User, bool) {
	var user *model.User
	database.DB.First(&user, "id = ?", id)

	if user.ID == 0 {
		return nil, false
	}

	return user, true
}

func InsertNew(user *model.User) error {
	err := database.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
