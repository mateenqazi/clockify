package users

import (
	"clockify/helpers"
	"clockify/models"
	"errors"
	"fmt"

	"clockify/types"

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

func (s *UserService) GetAllUser() ([]models.User, error) {
	var users []models.User

	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) RegisterUser(creds types.Credentials) (bool, error) {
	users := types.User{
		Email:    creds.Email,
		Password: helpers.HashPassword(creds.Password),
		IsActive: true,
	}

	if creds.Email == "" || creds.Password == "" {
		return false, errors.New("empty field are not allowed")
	}

	if !helpers.IsValidEmail(creds.Email) {
		return false, errors.New("email is not valid")
	}

	if ok, _, _ := helpers.IsEmailExists(s.db, creds.Email); !ok {
		result := s.db.Create(&users)
		if result.Error != nil {
			panic("Failed to save data into the database!")
		}

		fmt.Println("Data saved successfully!")
		return true, nil
	}

	fmt.Println("Email Already Exist")

	return false, nil
}

func (s *UserService) LoginUser(creds types.Credentials) (models.User, error) {
	emptyUser := models.User{}
	ok, result, _ := helpers.IsEmailExists(s.db, creds.Email)

	if ok {
		if !helpers.ComparePassword(creds.Password, result.Password) {
			fmt.Println("Password Does not Match")
			fmt.Println("Login Failed!")
			return emptyUser, errors.New("password does not matched")
		}
	}
	fmt.Println("Login Sucessfully!")
	return result, nil
}
