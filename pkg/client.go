package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"tripleqleads-demo/domain"
)

const (
	BaseURL          = "https://api.tripleqleads.com/v1"
	CompanyEndpoint  = "/enricher/company"
	EmployeeEndpoint = "/enricher/company/employees"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		apiKey:  apiKey,
		baseURL: BaseURL,
	}
}

func NewClientWithBaseURL(apiKey, baseURL string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

func (c *Client) GetCompany(req domain.EnrichmentRequest) (*domain.Company, error) {
	payload := map[string]string{}

	if req.CompanyLinkedInID != "" {
		payload["company_linkedin_id"] = req.CompanyLinkedInID
	} else if req.CompanyLinkedInURL != "" {
		payload["company_linkedin_url"] = req.CompanyLinkedInURL
	} else {
		return nil, fmt.Errorf("either company_linkedin_id or company_linkedin_url is required")
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.baseURL+CompanyEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-API-KEY", c.apiKey)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp domain.APIResponse
		if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Error != "" {
			return nil, fmt.Errorf("API error: %s", errorResp.Error)
		}
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var apiResp domain.CompanyAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if apiResp.Status != "OK" {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return &apiResp.Data.Company, nil
}

func (c *Client) GetEmployees(companyLinkedInID string, limit int) ([]domain.Employee, error) {
	payload := map[string]interface{}{
		"company_linkedin_id": companyLinkedInID,
		"limit":               limit,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.baseURL+EmployeeEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-API-KEY", c.apiKey)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp domain.APIResponse
		if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Error != "" {
			return nil, fmt.Errorf("API error: %s", errorResp.Error)
		}
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	var apiResp domain.EmployeeAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if apiResp.Status != "OK" {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	return apiResp.Data.Employees, nil
}
