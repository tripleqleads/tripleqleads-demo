package handlers

import (
	"net/http"
	"tripleqleads-demo/domain"
	"tripleqleads-demo/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	enrichmentService *services.EnrichmentService
}

func NewHandler(enrichmentService *services.EnrichmentService) *Handler {
	return &Handler{
		enrichmentService: enrichmentService,
	}
}

func (h *Handler) EnrichCompany(c *gin.Context) {
	var req domain.EnrichmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.APIResponse{
			Status: "ERROR",
			Error:  "Invalid request payload.",
		})
		return
	}

	if req.CompanyLinkedInID == "" && req.CompanyLinkedInURL == "" {
		c.JSON(http.StatusBadRequest, domain.APIResponse{
			Status: "ERROR",
			Error:  "company_linkedin_url or company_linkedin_id is required.",
		})
		return
	}

	result, err := h.enrichmentService.EnrichCompany(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status: "ERROR",
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.APIResponse{
		Status: "OK",
		Data:   result,
	})
}