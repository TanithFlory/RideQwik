package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rideqwik/api/internal/models"
	"github.com/rideqwik/api/internal/services"
)

type RideRequestHandler struct {
	rideService *services.RideService
}

func NewRideRequestHandler(rideService *services.RideService) *RideRequestHandler {
	return &RideRequestHandler{
		rideService: rideService,
	}
}

func (h *RideRequestHandler) RequestRides(c *gin.Context) {
	var req models.RideRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	response, err := h.rideService.GetRideOptions(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
