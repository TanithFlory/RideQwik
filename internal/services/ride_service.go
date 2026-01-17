package services

import (
	"context"
	"fmt"

	"github.com/rideqwik/api/internal/models"
)

type RideService struct {
	uberService *UberService
	// Add other providers here: lyftService, etc.
}

func NewRideService(uberService *UberService) *RideService {
	return &RideService{
		uberService: uberService,
	}
}

func (s *RideService) GetRideOptions(ctx context.Context, req *models.RideRequest) (*models.RideResponse, error) {
	var allOptions []models.RideOption

	if req.UberToken != "" {
		uberOptions, err := s.uberService.GetEstimates(ctx, req.UberToken, req)
		if err != nil {
			fmt.Printf("Uber API error: %v\n", err)
		} else {
			allOptions = append(allOptions, uberOptions...)
		}
	}

	if len(allOptions) == 0 {
		return nil, fmt.Errorf("no ride options available")
	}

	return &models.RideResponse{
		RequestID: generateRequestID(),
		Options:   allOptions,
	}, nil
}

func generateRequestID() string {
	return "req_" + "123456"
}
