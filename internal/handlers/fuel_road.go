package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllFuelRoad(c *gin.Context) {
	fuelRoads, err := h.Services.FuelRoad.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fuelRoads)

}

func (h *Handler) createFuelRoad(c *gin.Context) {

}

func (h *Handler) deleteFuelRoad(c *gin.Context) {

}
func (h *Handler) updateFuelRoad(c *gin.Context) {

}
