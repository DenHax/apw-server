package handler

import (
	"apw/internal/middleware"
	"apw/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

type Response struct {
	Name string `json:"name"`
}

func GetPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Name: "u name"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.Use(middleware.CORSMIddleware())
	api := router.Group("/api")
	{
		fuel_road := api.Group("/fuel-road")
		{
			fuel_road.GET("/", h.getAllFuelRoad)
		}
		fuel_type := api.Group("/fuel-type")
		{
			fuel_type.GET("/", h.getAllFuelType)
		}
		upload := api.Group("/upload")
		{
			upload.GET("/", h.getAllUpload)
			upload.POST("/:id", h.createUpload)
			upload.DELETE("/:id", h.deleteUpload)
		}
		subsystem := api.Group("/subsystem")
		{
			subsystem.GET("/", h.getAllSubsystem)
		}
		employee := api.Group("/employee")
		{
			employee.GET("/", h.getAllEmployee)
		}

	}

	return router
}
