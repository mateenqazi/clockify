package controller

import (
	"clockify/domain/respository"
	"clockify/domain/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ActivitiesController struct {
	activitiesService services.ActivitiesServicesInterface
}

func NewActivitiesController(db *gorm.DB) *ActivitiesController {
	activitiesRepo := respository.NewActivitiesDBRepository(db)
	activitiesService := services.NewActivitiesService(activitiesRepo)
	return &ActivitiesController{
		activitiesService: activitiesService,
	}
}

// RegisterRoutes registers the activities-related API routes.
func (c *ActivitiesController) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/activities", c.GetAllActivities).Methods(http.MethodGet)
	router.HandleFunc("/activities", c.CreateActivity).Methods(http.MethodPost)
	router.HandleFunc("/activities/{id}", c.DeleteActivity).Methods(http.MethodDelete)
	router.HandleFunc("/activities", c.UpdateActivityName).Methods(http.MethodPut)
	router.HandleFunc("/activities/duplicate", c.DuplicateActivity).Methods(http.MethodPost)
	router.HandleFunc("/activities/search", c.SearchActivities).Methods(http.MethodGet).Queries("query", "{query}", "userId", "{userId}")
}

func (c *ActivitiesController) GetAllActivities(w http.ResponseWriter, r *http.Request) {
	c.activitiesService.GetAllActitives(w, r)
}

func (c *ActivitiesController) CreateActivity(w http.ResponseWriter, r *http.Request) {
	c.activitiesService.CreateActivities(w, r)
}

func (c *ActivitiesController) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	activityId := mux.Vars(r)["id"]

	if activityId == "" {
		log.Println("delete failed")
		http.Error(w, "activity id is empty", http.StatusBadRequest)
		return
	}

	c.activitiesService.DeleteActivity(w, r)
}

func (c *ActivitiesController) UpdateActivityName(w http.ResponseWriter, r *http.Request) {
	c.activitiesService.UpdateName(w, r)
}

func (c *ActivitiesController) DuplicateActivity(w http.ResponseWriter, r *http.Request) {
	c.activitiesService.DuplicateActivity(w, r)
}

func (c *ActivitiesController) SearchActivities(w http.ResponseWriter, r *http.Request) {
	c.activitiesService.SearchActivities(w, r)
}
