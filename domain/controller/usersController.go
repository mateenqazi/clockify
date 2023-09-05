// controller.go

package controller

import (
	"net/http"

	"clockify/domain/respository"
	"clockify/domain/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserController struct {
	userService services.UserServicesInterface
}

func NewUserController(db *gorm.DB) *UserController {
	userRepo := respository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	return &UserController{
		userService: userService,
	}
}

func (c *UserController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", c.GetAllUser).Methods(http.MethodGet)
	router.HandleFunc("/users", c.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/users/login", c.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", c.DeleteUser).Methods(http.MethodDelete)
}

func (c *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	c.userService.GetAllUser(w, r)
}

func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	c.userService.RegisterUser(w, r)
}

func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	c.userService.LoginUser(w, r)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	c.userService.DeleteUser(w, r)
}
