package database

import (
	"praktikum/config"
	"praktikum/middleware"
	"praktikum/model"
)

func GetDetailUsers(userId int) (interface{}, error) {
	var user model.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *model.User) (interface{}, error) {
	var err error

	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	userResponse := model.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return userResponse, nil
}
