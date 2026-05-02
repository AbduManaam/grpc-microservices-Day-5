package main

import (
    "api-gateway/clients"
    "api-gateway/config"
    "api-gateway/handlers"
    "api-gateway/middleware"

    "github.com/gin-gonic/gin"
)

func main() {
	userAddr := config.GetEnv("USER_SERVICE_ADDR", "localhost:50051")
	orderAddr := config.GetEnv("ORDER_SERVICE_ADDR", "localhost:50052")
	clients.InitUserClient(userAddr)
	clients.InitOrderClient(orderAddr)

    r := gin.Default()

    // Public routes — no JWT needed
    r.POST("/register", handlers.Register)
    r.POST("/login",    handlers.Login)

    // Protected routes — JWT required
    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/user/:id", handlers.GetUser)
        protected.POST("/order",   handlers.CreateOrder)
    }

    port := config.GetEnv("PORT", "8080")
    r.Run(":" + port)
}