package main

import (
    "api-gateway/clients"
    "api-gateway/handlers"
    "api-gateway/middleware"

    "github.com/gin-gonic/gin"
)

func main() {
    clients.InitUserClient("user-service:50051")
    clients.InitOrderClient("order-service:50052")

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

    r.Run(":8080")
}