package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserSubrouter(r *mux.Router, db *gorm.DB) {

	userService := NewUserService(db)
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/create", userService.RegisterUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/delete/{id}", userService.DeleteUser).Methods(http.MethodDelete)
	userRouter.HandleFunc("/login", userService.LoginUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/get", userService.GetAllUser).Methods(http.MethodGet)
}
