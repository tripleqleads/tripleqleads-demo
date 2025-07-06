# TripleQLeads Demo API

A Go-based demo API server that enriches company data using the TripleQLeads API. This server provides a single endpoint that fetches both company information and employee data (limited to 5 employees) with built-in rate limiting.

## Features

- **Company Enrichment**: Fetch company data using LinkedIn URL or LinkedIn ID
- **Employee Data**: Automatically retrieves up to 5 employees for each company
- **Rate Limiting**: 1 request per IP per minute
- **Flexible Configuration**: Support for development and production environments
- **Built with Gin**: Fast HTTP framework with middleware support

## Quick Start

### Prerequisites

- Go 1.23.4 or later
- TripleQLeads API key (with `key_` prefix)

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd tripleqleads-demo

# Install dependencies
go mod tidy

# Build the server
go build -o server ./cmd/server
```

### Running the Server

#### Basic Usage
```bash
./server -api-key="key_your_actual_api_key_here"
```

#### Using Environment Variables
```bash
export TRIPLEQLEADS_API_KEY="key_your_api_key_here"
./server
```

#### Custom Port
```bash
./server -api-key="key_your_api_key_here" -port="3000"
```

### Configuration Options

| Flag | Environment Variable | Default | Description |
|------|---------------------|---------|-------------|
| `-api-key` | `TRIPLEQLEADS_API_KEY` | *(required)* | TripleQLeads API key |
| `-port` | `PORT` | `8080` | Server port |

*Note: Command line flags take priority over environment variables.*

## API Usage

### Endpoint

```
POST /v1/enricher/company
```

### Request Body

You can provide either a LinkedIn URL or LinkedIn ID:

```json
{
  "company_linkedin_url": "https://www.linkedin.com/company/microsoft/"
}
```

or

```json
{
  "company_linkedin_id": "1035"
}
```

### Example Requests

#### Using curl
```bash
# With LinkedIn URL
curl -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{
    "company_linkedin_url": "https://www.linkedin.com/company/microsoft/"
  }' | jq '.'

# With LinkedIn ID
curl -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{
    "company_linkedin_id": "1035"
  }' | jq '.'
```

### Response Format

#### Success Response (200 OK)
```json
{
  "status": "OK",
  "data": {
    "company": {
      "linkedin_id": "1035",
      "name": "Microsoft",
      "description": "...",
      "linkedin_url": "https://www.linkedin.com/company/microsoft/",
      "estimated_employee_count": 221000,
      "locations": [...],
      "industry": [...]
    },
    "employees": [
      {
        "linkedin_id": "...",
        "name": "John Doe",
        "headline": "Software Engineer at Microsoft",
        "current_position": {
          "company_data": {...},
          "role": "Software Engineer",
          "tenure_at_company": {
            "years": 3,
            "months": 6
          }
        }
      }
    ]
  }
}
```

#### Error Responses

**Rate Limited (429)**
```json
{
  "status": "ERROR",
  "error": "Rate limit exceeded. Only 1 request per IP is allowed."
}
```

**Bad Request (400)**
```json
{
  "status": "ERROR",
  "error": "company_linkedin_url or company_linkedin_id is required."
}
```

**Internal Server Error (500)**
```json
{
  "status": "ERROR",
  "error": "API error: Company couldn't be found with the provided identifier."
}
```

## Rate Limiting

- **Limit**: 1 request per IP per minute
- **Scope**: Per IP address
- **Response**: HTTP 429 when exceeded
- **Reset**: Automatically resets after 1 minute

## Project Structure

```
tripleqleads-demo/
├── cmd/server/           # Application entry point
│   └── main.go
├── domain/              # Data models and types
│   └── domain.go
├── pkg/                 # Shared packages
│   ├── client.go        # TripleQLeads API client
│   ├── middleware.go    # Rate limiting middleware
│   └── handlers/        # HTTP handlers
│       └── handlers.go
├── services/            # Business logic
│   └── enrichment.go
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
└── README.md           # This file
```

## Development

### Running Tests
```bash
go test ./...
```

### Code Formatting
```bash
go fmt ./...
```

### Linting
```bash
go vet ./...
```

### Building for Production
```bash
go build -ldflags="-s -w" -o server ./cmd/server
```

## Examples

### Test Different Companies
```bash
# Microsoft
curl -s -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{"company_linkedin_url": "https://www.linkedin.com/company/microsoft/"}' | jq '.'

# Google
curl -s -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{"company_linkedin_id": "1441"}' | jq '.'

# Apple
curl -s -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{"company_linkedin_url": "https://www.linkedin.com/company/apple/"}' | jq '.'
```

### Extract Specific Data
```bash
# Get only company name and employee count
curl -s -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{"company_linkedin_url": "https://www.linkedin.com/company/microsoft/"}' | \
  jq '.data.company | {name: .name, employee_count: .estimated_employee_count}'

# Get employee names only
curl -s -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{"company_linkedin_url": "https://www.linkedin.com/company/google/"}' | \
  jq '.data.employees[].name'
```

### Test Rate Limiting
```bash
# Run these commands in quick succession
curl -s -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{"company_linkedin_url": "https://www.linkedin.com/company/netflix/"}' | jq '.'

# This should return a 429 error
curl -s -X POST http://localhost:8080/v1/enricher/company \
  -H "Content-Type: application/json" \
  -d '{"company_linkedin_url": "https://www.linkedin.com/company/netflix/"}' | jq '.'
```

## Troubleshooting

### Common Issues

**"API key is required" Error**
- Ensure you've set the API key via flag or environment variable
- Check that your API key has the correct `key_` prefix

**Rate Limit Exceeded**
- Wait 1 minute between requests from the same IP
- Consider implementing client-side rate limiting

**Connection Refused**
- Verify the base URL is correct for your environment
- Check that the TripleQLeads API is accessible

**Invalid Company Identifier**
- Ensure LinkedIn URLs are in the correct format
- Verify the company exists on LinkedIn

### Debug Mode
```bash
# Run with debug logging
GIN_MODE=debug ./server -api-key="your_key"
```

## Docker Deployment

### Building the Docker Image

```bash
# Build the image
docker build -t tripleqleads-demo .

# Or with a specific tag
docker build -t tripleqleads-demo:v1.0.0 .
```

### Running with Docker

#### Basic Usage
```bash
docker run -d \
  --name tripleqleads-demo \
  -p 8080:8080 \
  -e TRIPLEQLEADS_API_KEY="key_your_actual_api_key_here" \
  tripleqleads-demo
```

#### Custom Port
```bash
docker run -d \
  --name tripleqleads-demo \
  -p 3000:3000 \
  -e TRIPLEQLEADS_API_KEY="key_your_actual_api_key_here" \
  -e PORT="3000" \
  tripleqleads-demo
```

#### Using Environment File
Create a `.env` file:
```bash
# .env
TRIPLEQLEADS_API_KEY=key_your_actual_api_key_here
PORT=8080
```

Run with environment file:
```bash
docker run -d \
  --name tripleqleads-demo \
  -p 8080:8080 \
  --env-file .env \
  tripleqleads-demo
```

### Docker Compose

Create a `docker-compose.yml` file:
```yaml
version: '3.8'

services:
  tripleqleads-demo:
    build: .
    ports:
      - "8080:8080"
    environment:
      - TRIPLEQLEADS_API_KEY=key_your_actual_api_key_here
      - PORT=8080
    restart: unless-stopped
```

Run with Docker Compose:
```bash
# Start the service
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the service
docker-compose down
```

### Docker Commands

```bash
# View running containers
docker ps

# View logs
docker logs tripleqleads-demo

# Follow logs
docker logs -f tripleqleads-demo

# Stop container
docker stop tripleqleads-demo

# Remove container
docker rm tripleqleads-demo

# Remove image
docker rmi tripleqleads-demo
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is for demonstration purposes.