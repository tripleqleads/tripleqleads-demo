# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based demo application for the TripleQLeads API using Go 1.23.4. The project implements a REST API for company and employee data enrichment based on LinkedIn information.

## Project Structure

```
tripleqleads-demo/
├── cmd/server/           # Application entrypoint (main.go)
├── domain/              # Business domain logic and entities
├── pkg/                 # Shared packages and utilities
├── services/            # Application services and business logic
├── api_docs/            # API documentation
│   ├── 1.getting-started/
│   └── 2.enricher/
└── go.mod               # Go module definition
```

## API Overview

The TripleQLeads API provides two main endpoints:

### Company Enrichment
- **Endpoint**: `POST https://api.tripleqleads.com/v1/enricher/company`
- **Purpose**: Enrich company data using LinkedIn URL or Company LinkedIn ID
- **Authentication**: X-API-KEY header required
- **Cost**: 1 daily credit per request

### Employee Data
- **Endpoint**: `POST https://api.tripleqleads.com/v1/enricher/company/employees`
- **Purpose**: Fetch employees at a company using Company LinkedIn ID
- **Authentication**: X-API-KEY header required
- **Pagination**: Uses cursor-based pagination for large datasets (limit 1-100)

## Common Development Commands

### Build
```bash
go build ./cmd/server
```

### Run
```bash
go run ./cmd/server
```

### Test
```bash
go test ./...
```

### Format Code
```bash
go fmt ./...
```

### Lint
```bash
go vet ./...
```

### Get Dependencies
```bash
go mod tidy
```

## Key API Data Models

### Company Data
- Contains LinkedIn metadata, company details, locations, employee counts, and industry information
- Includes nested Location objects with address and HQ information

### Employee Data
- Contains LinkedIn profile information, current position details, and tenure information
- Includes nested Company Data and Position objects

## Authentication & Error Handling

- API uses X-API-KEY header authentication with `key_` prefix
- Standard HTTP status codes (200, 400, 401, 403, 404, 429, 500)
- Structured error responses with status and error message fields
- Rate limiting and daily credit consumption tracking

## Development Notes

- The project appears to be in early development stages with minimal Go implementation
- API documentation is comprehensive and follows OpenAPI-style conventions
- The demo will likely implement HTTP handlers that mirror the documented API endpoints
- Consider implementing rate limiting, authentication middleware, and structured logging