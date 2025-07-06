package services

import (
	"fmt"
	"tripleqleads-demo/domain"
	"tripleqleads-demo/pkg"
)

type EnrichmentService struct {
	client *pkg.Client
}

func NewEnrichmentService(apiKey string) *EnrichmentService {
	return &EnrichmentService{
		client: pkg.NewClient(apiKey),
	}
}

func NewEnrichmentServiceWithBaseURL(apiKey, baseURL string) *EnrichmentService {
	return &EnrichmentService{
		client: pkg.NewClientWithBaseURL(apiKey, baseURL),
	}
}

func (s *EnrichmentService) EnrichCompany(req domain.EnrichmentRequest) (*domain.EnrichmentResponse, error) {
	company, err := s.client.GetCompany(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get company data: %w", err)
	}

	employees, err := s.client.GetEmployees(company.LinkedInID, 5)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee data: %w", err)
	}

	return &domain.EnrichmentResponse{
		Company:   *company,
		Employees: employees,
	}, nil
}