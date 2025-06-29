.PHONY: run build test clean

# Run the application
run:
	go run cmd/server/main.go

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Test the application
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run with hot reload (requires air)
dev:
	air

# Database migration (if using migrate tool)
migrate-up:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/go_clean_db?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/go_clean_db?sslmode=disable" down