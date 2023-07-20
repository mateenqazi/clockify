package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserSubrouter(r *mux.Router, db *gorm.DB) {

	userService := NewUserService(db)
	userRouter := r.PathPrefix("/users").Subrouter()
	// userRouter.HandleFunc("/create", createUserHandler).Methods(http.MethodPost)
	// userRouter.HandleFunc("/delete", createUserHandler).Methods(http.MethodDelete)
	// userRouter.HandleFunc("/login", userService.GetAllUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/get", userService.GetAllUser).Methods(http.MethodGet)
}
