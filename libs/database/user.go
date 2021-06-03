package database

import (
	"acp8/configs"
	"acp8/models"
)

func GetUsers() (interface{}, error) {

	var dataUser []models.User

	err := configs.DB.Find(&dataUser).Error

	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

func RegisterUser(userRegister models.UserRegister) (interface{}, error) {

	var userDatabase models.User = models.User{}
	userDatabase.Name = userRegister.Name
	userDatabase.Password = userRegister.Password
	userDatabase.Email = userRegister.Email

	err := configs.DB.Create(&userDatabase).Error

	if err != nil {
		return nil, err
	}

	return userDatabase, nil
}
