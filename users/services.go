package users

import (
	"clockify/helpers"
	"clockify/models"
	"errors"
	"log"
	"net/http"

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

func (s *UserService) GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := s.db.Model(&users).Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.SendJSONResponse(w, http.StatusOK, users)

}

func (s *UserService) RegisterUser(creds types.Credentials) (bool, error) {
	users := models.User{
		Email:    creds.Email,
		IsActive: true,
	}

	hashPassword, err := helpers.HashPassword(creds.Password)
	if err != nil {
		return false, errors.New("something went wrong")
	}

	users.Password = hashPassword

	if creds.Email == "" || creds.Password == "" {
		return false, errors.New("empty field are not allowed")
	}

	if !helpers.IsValidEmail(creds.Email) {
		return false, errors.New("email is not valid")
	}

	if ok, _, _ := users.IsEmailExists(s.db, creds.Email); !ok {
		result := s.db.Model(&users).Create(&users)
		if result.Error != nil {
			panic("Failed to save data into the database!")
		}

		log.Println("Data saved successfully!")

		return true, nil
	}

	log.Println("Email Already Exist")

	return false, nil
}

func (s *UserService) LoginUser(creds types.Credentials) (models.User, error) {
	emptyUser := models.User{}

	ok, result, _ := emptyUser.IsEmailExists(s.db, creds.Email)
	if !ok {
		return emptyUser, errors.New("email not exists")
	}

	if err := helpers.IsPasswordMatch(creds.Password, result.Password); err != nil {
		log.Fatal(err)
		return emptyUser, errors.New("password does not matched")
	}

	log.Println("Login Sucessfully!", result)

	return result, nil
}

func (s *UserService) DeleteUser(UserId int) (bool, error) {
	var user models.User
	if err := s.db.Model(&user).Where("id = ?", UserId).Delete(&user); err.Error != nil {
		log.Fatal("delete failed")
		return false, errors.New("delete failed")
	}

	return true, nil
}
