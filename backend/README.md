# Go Clean Architecture API

Simple REST API built with Go, Gin, GORM, and PostgreSQL following Clean Architecture principles.

## Project Structure

```
go-clean/
├── cmd/api/                 # Application entry point
├── internal/
│   ├── domain/             # Business entities and repository interfaces
│   ├── usecase/            # Business logic layer
│   ├── infrastructure/     # External dependencies (database)
│   └── delivery/http/      # HTTP handlers and routing
├── pkg/config/             # Configuration
└── .env.example           # Environment variables example
```

## Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your database settings
3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Install air for hot reload (optional):
   ```bash
   go install github.com/cosmtrek/air@v1.43.0
   ```

5. Run the application:
   ```bash
   # Regular run
   go run cmd/api/main.go
   
   # With hot reload (recommended for development)
   air
   ```

## API Endpoints

### Users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Example Request

Create User:
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

## Database Setup

Make sure PostgreSQL is running and create a database named `goclean` (or whatever you set in DB_NAME).