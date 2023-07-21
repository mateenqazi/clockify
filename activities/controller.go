package activities

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserSubrouter(r *mux.Router, db *gorm.DB) {

	activitiesService := NewActivitiesService(db)

	activitiesRouter := r.PathPrefix("/activities").Subrouter()

	// apis
	activitiesRouter.HandleFunc("/create", activitiesService.CreateActivities).Methods(http.MethodPost)
	activitiesRouter.HandleFunc("/delete/{id}", activitiesService.DeleteActivity).Methods(http.MethodDelete)
	activitiesRouter.HandleFunc("/search", activitiesService.SearchActivities).Methods(http.MethodGet)
	activitiesRouter.HandleFunc("/update", activitiesService.UpdateName).Methods(http.MethodPut)
	activitiesRouter.HandleFunc("/copy", activitiesService.DuplicateActivity).Methods(http.MethodGet)
	activitiesRouter.HandleFunc("/get", activitiesService.GetAllActitives).Methods(http.MethodGet)
}
