package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rideqwik/api/internal/config"
	"github.com/rideqwik/api/internal/models"
)

type UberService struct {
	config *config.Config
	client *http.Client
}

type UberTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func NewUberService(cfg *config.Config) *UberService {
	return &UberService{
		config: cfg,
		client: &http.Client{},
	}
}

// ExchangeCodeForToken exchanges authorization code for access token
// Frontend will store this token
func (s *UberService) ExchangeCodeForToken(ctx context.Context, code string) (*UberTokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", s.config.UberClientID)
	data.Set("client_secret", s.config.UberClientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", s.config.UberRedirectURI)
	data.Set("code", code)
	data.Set("scope", "profile request")

	req, err := http.NewRequestWithContext(ctx, "POST", "https://auth.uber.com/oauth/v2/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token exchange request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("token exchange failed: status %d, body: %s", resp.StatusCode, string(body))
	}

	var tokenResp UberTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	return &tokenResp, nil
}

// GetEstimates gets ride estimates using token from frontend
func (s *UberService) GetEstimates(ctx context.Context, uberToken string, req *models.RideRequest) ([]models.RideOption, error) {
	if uberToken == "" {
		return nil, fmt.Errorf("uber token is required")
	}

	apiURL := fmt.Sprintf(
		"%s/v1.2/estimates/price?start_latitude=%f&start_longitude=%f&end_latitude=%f&end_longitude=%f",
		s.config.UberAPIBaseURL,
		req.PickupLatitude,
		req.PickupLongitude,
		req.DropoffLatitude,
		req.DropoffLongitude,
	)

	httpReq, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+uberToken)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("uber api request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("uber api error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var uberResp UberPriceEstimateResponse
	if err := json.NewDecoder(resp.Body).Decode(&uberResp); err != nil {
		return nil, fmt.Errorf("failed to decode uber response: %w", err)
	}

	options := make([]models.RideOption, 0, len(uberResp.Prices))
	for _, price := range uberResp.Prices {
		options = append(options, models.RideOption{
			Platform:      "uber",
			ProductID:     price.ProductID,
			DisplayName:   price.DisplayName,
			EstimatedFare: (price.LowEstimate + price.HighEstimate) / 2,
			ETAMinutes:    price.Duration / 60,
			Currency:      price.CurrencyCode,
		})
	}

	return options, nil
}

type UberPriceEstimateResponse struct {
	Prices []UberPrice `json:"prices"`
}

type UberPrice struct {
	ProductID    string  `json:"product_id"`
	DisplayName  string  `json:"display_name"`
	LowEstimate  float64 `json:"low_estimate"`
	HighEstimate float64 `json:"high_estimate"`
	CurrencyCode string  `json:"currency_code"`
	Duration     int     `json:"duration"`
}
