package main

import (
	"log"
	"time"

	"go-clean/internal/delivery/http"
	adminHandler "go-clean/internal/delivery/http/admin"
	authHandler "go-clean/internal/delivery/http/auth"
	orderHandler "go-clean/internal/delivery/http/order"
	productHandler "go-clean/internal/delivery/http/product"
	userHandler "go-clean/internal/delivery/http/user"
	"go-clean/internal/domain/order"
	"go-clean/internal/domain/product"
	"go-clean/internal/domain/user"
	orderRepo "go-clean/internal/repository/order"
	productRepo "go-clean/internal/repository/product"
	userRepo "go-clean/internal/repository/user"
	authUseCase "go-clean/internal/usecase/auth"
	orderUseCase "go-clean/internal/usecase/order"
	productUseCase "go-clean/internal/usecase/product"
	userUseCase "go-clean/internal/usecase/user"
	"go-clean/pkg/config"
	"go-clean/pkg/database"
	"go-clean/pkg/jwt"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&user.User{}, &product.Product{}, &order.Order{}, &order.OrderItem{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	userRepository := userRepo.NewPostgresRepository(db)
	productRepository := productRepo.NewPostgresRepository(db)
	orderRepository := orderRepo.NewPostgresRepository(db)

	jwtManager := jwt.NewJWTManager(cfg.JWTSecret, "go-clean-api", 24*time.Hour, 7*24*time.Hour)

	userService := userUseCase.NewService(userRepository)
	productService := productUseCase.NewService(productRepository)
	orderService := orderUseCase.NewService(orderRepository, productRepository, userRepository)
	authService := authUseCase.NewService(userRepository, jwtManager)

	authH := authHandler.NewHandler(authService)
	userH := userHandler.NewHandler(userService)
	productH := productHandler.NewHandler(productService)
	orderH := orderHandler.NewHandler(orderService)
	adminH := adminHandler.NewHandler(userService)

	router := http.SetupRoutes(authH, userH, productH, orderH, adminH, jwtManager)

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}