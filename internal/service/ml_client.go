package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AIdenTale/go-hack.git/internal/model"
)

type MLClient struct {
	baseURL    string
	httpClient *http.Client
}

type mlRequestData struct {
	FHRTime   []float64 `json:"fhr_time"`
	FHRValues []float64 `json:"fhr_values"`
	UCTime    []float64 `json:"uc_time"`
	UCValues  []float64 `json:"uc_values"`
}

func NewMLClient(baseURL string) *MLClient {
	return &MLClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *MLClient) Predict(data *mlRequestData) (*model.MLPredictResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(c.baseURL+"/predict", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("predict request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("predict request failed: %d", resp.StatusCode)
	}

	var result model.MLPredictResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode predict response: %w", err)
	}

	return &result, nil
}

func (c *MLClient) GetFeatures(data *mlRequestData) (*model.MLFeaturesResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(c.baseURL+"/features", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("features request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("features request failed: %d", resp.StatusCode)
	}

	var result model.MLFeaturesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode features response: %w", err)
	}

	return &result, nil
}