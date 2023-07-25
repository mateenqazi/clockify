package services

import (
	"clockify/helpers"
	"encoding/json"
	"log"
	"net/http"

	"clockify/domain/entity"
	"clockify/domain/respository"

	"github.com/gorilla/mux"
)

type UserServicesInterface interface {
	GetAllUser(w http.ResponseWriter, r *http.Request)
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type UserService struct {
	userRepository respository.UserRepository
}

func NewUserService(userRepository respository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := s.userRepository.GetAllUser()
	if err != nil {
		log.Println(err.Error())
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

	users := entity.User{
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

	isExist, _ := s.userRepository.FindByEmail(email)
	if isExist == nil {
		result, err := s.userRepository.Create(&users)
		if err != nil {
			http.Error(w, "Failed to save data into the database!", http.StatusBadRequest)
			return
		}

		log.Println("Data saved successfully!")

		helpers.SendJSONResponse(w, http.StatusOK, result)
		return
	}

	log.Println("Email Already Exists", isExist)

	http.Error(w, "Email Already Exists", http.StatusBadRequest)
}

func (s *UserService) LoginUser(w http.ResponseWriter, r *http.Request) {
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

	userValue, _ := s.userRepository.FindByEmail(email)

	if userValue == nil {
		http.Error(w, "email don't exist", http.StatusBadRequest)
		return
	}

	if err := helpers.IsPasswordMatch(password, userValue.Password); err != nil {
		log.Println(err)
		http.Error(w, "password does not matched", http.StatusBadRequest)
		return
	}

	log.Println("Login Sucessfully!", userValue)

	helpers.SendJSONResponse(w, http.StatusOK, userValue)
}

func (s *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	if userId == "" {
		log.Println("delete failed")
		http.Error(w, "user id is empty", http.StatusBadRequest)
		return
	}

	if err := s.userRepository.Delete(userId); err != nil {
		log.Println("delete failed")
		http.Error(w, "user id is empty", http.StatusBadRequest)
		return
	}

	log.Println("user delete sucessfully!")

	helpers.SendJSONResponse(w, http.StatusNoContent, nil)
}
