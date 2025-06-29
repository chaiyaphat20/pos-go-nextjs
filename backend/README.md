# Go Clean Architecture Backend

A modern Go backend application built with Clean Architecture principles, designed for scalability and maintainability.

## 🏗️ Architecture Overview

This project follows **Clean Architecture** (also known as Hexagonal Architecture) with clear separation of concerns:

```
┌─────────────────────────────────────────────┐
│                 Delivery Layer              │  ← HTTP Handlers, REST API
├─────────────────────────────────────────────┤
│                Use Case Layer               │  ← Business Logic, Application Services
├─────────────────────────────────────────────┤
│                Domain Layer                 │  ← Entities, Repository Interfaces
├─────────────────────────────────────────────┤
│              Infrastructure Layer           │  ← Database, External Services
└─────────────────────────────────────────────┘
```

## 📂 Project Structure

```
backend/
├── cmd/server/main.go              # Application entry point
├── internal/                       # Private application code
│   ├── domain/                     # 🔥 Domain Layer (Entities & Interfaces)
│   │   ├── user/
│   │   │   ├── model.go           # User entity with business rules
│   │   │   └── repository.go      # User repository interface
│   │   ├── product/
│   │   │   ├── model.go           # Product entity with stock management
│   │   │   └── repository.go      # Product repository interface
│   │   └── order/
│   │       ├── model.go           # Order/OrderItem entities
│   │       └── repository.go      # Order repository interface
│   ├── usecase/                   # 📋 Use Case Layer (Business Logic)
│   │   ├── user/
│   │   │   ├── usecase.go         # User business logic interface
│   │   │   └── service.go         # User business logic implementation
│   │   ├── product/
│   │   │   ├── usecase.go         # Product business logic interface
│   │   │   └── service.go         # Product business logic implementation
│   │   └── order/
│   │       ├── usecase.go         # Order business logic interface
│   │       └── service.go         # Order business logic implementation
│   ├── repository/                # 🗄️ Infrastructure Layer (Data Access)
│   │   ├── user/postgres.go       # User PostgreSQL implementation
│   │   ├── product/postgres.go    # Product PostgreSQL implementation
│   │   └── order/postgres.go      # Order PostgreSQL implementation
│   └── delivery/http/             # 🌐 Delivery Layer (HTTP Handlers)
│       ├── router.go              # Route definitions
│       ├── user/handler.go        # User HTTP endpoints
│       ├── product/handler.go     # Product HTTP endpoints
│       └── order/handler.go       # Order HTTP endpoints
├── pkg/                           # 📦 Shared packages (public)
│   ├── config/config.go           # Configuration management
│   └── database/postgres.go       # Database connection
├── go.mod                         # Go module dependencies
├── go.sum                         # Dependency checksums
├── Makefile                       # Build and development commands
└── .air.toml                      # Hot reload configuration
```

## 🔧 Features

### User Management
- ✅ User registration and authentication
- ✅ Role-based access (Admin, User, Moderator)
- ✅ Password hashing with bcrypt
- ✅ User CRUD operations

### Product Management
- ✅ Product catalog with CRUD operations
- ✅ Stock management and tracking
- ✅ Low stock alerts
- ✅ Product search functionality
- ✅ Price and inventory management

### Order Management
- ✅ Order creation with multiple items
- ✅ Automatic stock deduction
- ✅ Order status tracking (Pending → Confirmed → Shipped → Delivered)
- ✅ Order cancellation with stock restoration
- ✅ User order history

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- PostgreSQL 12 or higher
- Air (for hot reload development)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd backend
   ```

2. **Install dependencies**
   ```bash
   make deps
   ```

3. **Set up database**
   ```bash
   # Create PostgreSQL database
   createdb goclean
   ```

4. **Configure environment** (optional)
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASS=password
   export DB_NAME=goclean
   export SERVER_PORT=8080
   ```

### Running the Application

**Development (with hot reload):**
```bash
make dev
# or
air
```

**Production:**
```bash
make run
# or
go run cmd/server/main.go
```

**Build executable:**
```bash
make build
./bin/server
```

## 📡 API Endpoints

### Users API
```
POST   /api/v1/users           # Create user
GET    /api/v1/users           # List users (with pagination)
GET    /api/v1/users/:id       # Get user by ID
PUT    /api/v1/users/:id       # Update user
DELETE /api/v1/users/:id       # Delete user
```

### Products API
```
POST   /api/v1/products           # Create product
GET    /api/v1/products           # List products (with search & pagination)
GET    /api/v1/products/:id       # Get product by ID
PUT    /api/v1/products/:id       # Update product
DELETE /api/v1/products/:id       # Delete product
PATCH  /api/v1/products/:id/stock # Update product stock
GET    /api/v1/products/low-stock # Get low stock products
```

### Orders API
```
POST   /api/v1/orders           # Create order
GET    /api/v1/orders           # List orders (with filters)
GET    /api/v1/orders/:id       # Get order by ID
PATCH  /api/v1/orders/:id/status # Update order status
POST   /api/v1/orders/:id/cancel # Cancel order
```

### Query Parameters
- `limit` - Pagination limit (default: 10)
- `offset` - Pagination offset (default: 0)
- `search` - Search products by name
- `user_id` - Filter orders by user
- `status` - Filter orders by status
- `threshold` - Low stock threshold (default: 10)

## 🏛️ Clean Architecture Benefits

### 1. **Dependency Inversion**
- Domain layer doesn't depend on external frameworks
- Use cases define interfaces, infrastructure implements them
- Easy to swap databases or external services

### 2. **Testability**
- Each layer can be tested in isolation
- Mock implementations for external dependencies
- Business logic is framework-independent

### 3. **Maintainability**
- Clear separation of concerns
- Single responsibility principle
- Easy to locate and modify specific functionality

### 4. **Scalability**
- Modular structure allows team scaling
- Independent deployment of services
- Easy to add new features without affecting existing code

## 🛠️ Development Commands

```bash
# Run with hot reload
make dev

# Run normally
make run

# Build executable
make build

# Install dependencies
make deps

# Run tests
make test

# Clean build artifacts
make clean
```

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | localhost | Database host |
| `DB_PORT` | 5432 | Database port |
| `DB_USER` | postgres | Database username |
| `DB_PASS` | password | Database password |
| `DB_NAME` | goclean | Database name |
| `SERVER_PORT` | 8080 | Server port |

## 📝 Example Requests

### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123",
    "role": "user"
  }'
```

### Create Product
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Gaming Laptop",
    "description": "High-performance gaming laptop",
    "price": 1299.99,
    "stock": 50
  }'
```

### Create Order
```bash
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "123e4567-e89b-12d3-a456-426614174000",
    "order_items": [
      {
        "product_id": "456e7890-e89b-12d3-a456-426614174001",
        "quantity": 2
      }
    ]
  }'
```

## 🤝 Contributing

1. Follow the Clean Architecture principles
2. Add tests for new features
3. Update documentation
4. Follow Go naming conventions
5. Use meaningful commit messages

## 📄 License

This project is licensed under the MIT License.