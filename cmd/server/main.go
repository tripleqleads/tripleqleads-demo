package main

import (
	"flag"
	"log"
	"os"
	"time"
	"tripleqleads-demo/pkg"
	"tripleqleads-demo/pkg/handlers"
	"tripleqleads-demo/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var apiKeyFlag = flag.String("api-key", "", "TripleQLeads API key")
	var portFlag = flag.String("port", "8080", "Port to run the server on")
	var baseURLFlag = flag.String("base-url", "", "Base URL for TripleQLeads API (default: https://api.tripleqleads.com/v1)")
	flag.Parse()

	apiKey := *apiKeyFlag
	if apiKey == "" {
		apiKey = os.Getenv("TRIPLEQLEADS_API_KEY")
	}
	if apiKey == "" {
		log.Fatal("API key is required. Set TRIPLEQLEADS_API_KEY environment variable or use -api-key flag")
	}

	baseURL := *baseURLFlag
	if baseURL == "" {
		baseURL = os.Getenv("TRIPLEQLEADS_BASE_URL")
	}

	var enrichmentService *services.EnrichmentService
	if baseURL != "" {
		enrichmentService = services.NewEnrichmentServiceWithBaseURL(apiKey, baseURL)
		log.Printf("Using custom base URL: %s", baseURL)
	} else {
		enrichmentService = services.NewEnrichmentService(apiKey)
	}
	handler := handlers.NewHandler(enrichmentService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"Accept",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(pkg.NewRateLimitMiddleware())

	r.POST("/v1/enricher/company", handler.EnrichCompany)

	port := *portFlag
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
