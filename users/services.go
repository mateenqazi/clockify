package users

import (
	"clockify/helpers"
	"clockify/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func (s *UserService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := data["email"].(string)
	password := data["password"].(string)

	users := models.User{
		Email:    email,
		IsActive: true,
	}

	hashPassword, err := helpers.HashPassword(password)
	if err != nil {
		log.Println(err)
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return
	}

	users.Password = hashPassword

	if email == "" || password == "" {
		http.Error(w, "empty fields are not allowed", http.StatusBadRequest)
		return
	}

	if !helpers.IsValidEmail(email) {
		http.Error(w, "email is not valid", http.StatusBadRequest)
		return
	}

	if ok, _, _ := users.IsEmailExists(s.db, email); !ok {
		result := s.db.Model(&users).Create(&users)
		if result.Error != nil {
			http.Error(w, "Failed to save data into the database!", http.StatusBadRequest)
			return
		}

		log.Println("Data saved successfully!")

		helpers.SendJSONResponse(w, http.StatusOK, result)
		return
	}

	log.Println("Email Already Exists")

	http.Error(w, "Email Already Exists", http.StatusBadRequest)
}

func (s *UserService) LoginUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := data["email"].(string)
	password := data["password"].(string)

	if email == "" || password == "" {
		http.Error(w, "empty fields are not allowed", http.StatusBadRequest)
		return
	}

	ok, result, _ := user.IsEmailExists(s.db, email)
	if !ok {
		http.Error(w, "email don't exist", http.StatusBadRequest)
		return
	}

	if err := helpers.IsPasswordMatch(password, result.Password); err != nil {
		log.Println(err)
		http.Error(w, "password does not matched", http.StatusBadRequest)
		return
	}

	log.Println("Login Sucessfully!", result)

	helpers.SendJSONResponse(w, http.StatusOK, result)
}

func (s *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userId := mux.Vars(r)["id"]

	if userId == "" {
		log.Println("delete failed")
		http.Error(w, "user id is empty", http.StatusBadRequest)
		return
	}

	if err := s.db.Model(&user).Where("id = ?", userId).Delete(&user); err.Error != nil {
		log.Println("delete failed")
		http.Error(w, "user id is empty", http.StatusBadRequest)
		return
	}

	log.Println("user delete sucessfully!")

	helpers.SendJSONResponse(w, http.StatusNoContent, nil)
}
