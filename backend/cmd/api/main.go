package main

import (
	"go-clean/internal/delivery/http"
	"go-clean/internal/infrastructure/database"
	"go-clean/internal/service"
	"go-clean/internal/usecase"
	"go-clean/pkg/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	userRepo := database.NewUserRepository(db)
	productRepo := database.NewProductRepository(db)
	saleRepo := database.NewSaleRepository(db)
	saleItemRepo := database.NewSaleItemRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)
	saleUsecase := usecase.NewSaleUsecase(saleRepo, saleItemRepo, productRepo)
	
	// JWT secret - ในโปรดักชันควรใช้จาก environment variable
	jwtSecret := "your-secret-key-here"
	authService := service.NewAuthService(userRepo, jwtSecret)

	userHandler := http.NewUserHandler(userUsecase)
	productHandler := http.NewProductHandler(productUsecase)
	saleHandler := http.NewSaleHandler(saleUsecase)
	authHandler := http.NewAuthHandler(authService)

	router := http.SetupRoutes(userHandler, productHandler, saleHandler, authHandler)

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}