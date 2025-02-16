package services

import (
	"bouguette/cmd/api/requests"
	"bouguette/common"
	"bouguette/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (userService UserService) RegisterUser(userRequest *requests.RegisterUserRequest) (*models.UserModel, error) {
	hashedPassword, err := common.HashPassword(userRequest.Password)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Registration failed")
	}

	createdUser := models.UserModel{
		FirstName: &userRequest.FirstName,
		LastName:  &userRequest.LastName,
		Email:     userRequest.Email,
		Password:  hashedPassword,
	}

	result := userService.db.Create(&createdUser)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, errors.New("Registration failed")
	}

	return &createdUser, nil
}

func (userService UserService) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	result := userService.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
