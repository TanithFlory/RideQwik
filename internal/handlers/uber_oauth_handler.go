package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rideqwik/api/internal/config"
	"github.com/rideqwik/api/internal/services"
)

type UberOAuthHandler struct {
	config      *config.Config
	uberService *services.UberService
}

func NewUberOAuthHandler(cfg *config.Config, uberService *services.UberService) *UberOAuthHandler {
	return &UberOAuthHandler{
		config:      cfg,
		uberService: uberService,
	}
}

func (h *UberOAuthHandler) AuthorizeUber(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "user not authenticated",
		})
		return
	}

	authURL := fmt.Sprintf(
		"https://auth.uber.com/oauth/v2/authorize?client_id=%s&response_type=code&scope=profile request&state=%d&redirect_uri=%s",
		h.config.UberClientID,
		userID.(int),
		h.config.UberRedirectURI,
	)

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"auth_url":    authURL,
		"redirect_to": authURL,
	})
}

func (h *UberOAuthHandler) UberCallback(c *gin.Context) {
	code := c.Query("code")
	errorParam := c.Query("error")

	if errorParam != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   fmt.Sprintf("uber authorization failed: %s", errorParam),
		})
		return
	}

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "authorization code not provided",
		})
		return
	}

	token, err := h.uberService.ExchangeCodeForToken(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   fmt.Sprintf("failed to exchange code for token: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "uber account connected successfully",
		"data":    token,
	})
}
