package http

import (
	adminHandler "go-clean/internal/delivery/http/admin"
	authHandler "go-clean/internal/delivery/http/auth"
	"go-clean/internal/delivery/http/middleware"
	orderHandler "go-clean/internal/delivery/http/order"
	productHandler "go-clean/internal/delivery/http/product"
	userHandler "go-clean/internal/delivery/http/user"
	"go-clean/pkg/jwt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(authH *authHandler.Handler, userH *userHandler.Handler, productH *productHandler.Handler, orderH *orderHandler.Handler, adminH *adminHandler.Handler, jwtManager *jwt.JWTManager) *gin.Engine {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/api/v1")
	{
		// Public auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authH.Login)
			auth.POST("/register", authH.Register)
			auth.POST("/refresh", authH.RefreshToken)
		}

		// Public product routes (anyone can view products)
		products := api.Group("/products")
		{
			products.GET("", productH.GetProducts)
			products.GET("/:id", productH.GetProduct)
		}

		// Public user creation route (for admin creating users without auth)
		api.POST("/users", userH.CreateUser)

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(jwtManager))
		{
			users := protected.Group("/users")
			{
				users.GET("", userH.GetUsers)
				users.GET("/:id", userH.GetUser)
				users.PUT("/:id", userH.UpdateUser)
				users.DELETE("/:id", userH.DeleteUser)
			}

			protectedProducts := protected.Group("/products")
			{
				protectedProducts.POST("", middleware.AdminOnly(), productH.CreateProduct)
				protectedProducts.PUT("/:id", middleware.AdminOnly(), productH.UpdateProduct)
				protectedProducts.DELETE("/:id", middleware.AdminOnly(), productH.DeleteProduct)
				protectedProducts.PATCH("/:id/stock", middleware.AdminOnly(), productH.UpdateStock)
				protectedProducts.GET("/low-stock", middleware.AdminOnly(), productH.GetLowStockProducts)
			}

			orders := protected.Group("/orders")
			{
				orders.POST("", orderH.CreateOrder)
				orders.GET("", orderH.GetOrders)
				orders.GET("/:id", orderH.GetOrder)
				orders.PATCH("/:id/status", middleware.AdminOnly(), orderH.UpdateOrderStatus)
				orders.POST("/:id/cancel", orderH.CancelOrder)
			}

			// Admin-only routes
			admin := protected.Group("/admin")
			admin.Use(middleware.AdminOnly())
			{
				admin.GET("/stats", adminH.GetSystemStats)
				admin.GET("/users/detailed", adminH.GetAllUsersDetailed)
			}
		}
	}

	return r
}