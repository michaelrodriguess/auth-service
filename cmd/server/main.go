package main

import (
	"context"
	"log"

	"github.com/michaelrodriguess/auth_service/config"
	"github.com/michaelrodriguess/auth_service/internal/handler"
	"github.com/michaelrodriguess/auth_service/internal/middleware"
	"github.com/michaelrodriguess/auth_service/internal/repository"
	"github.com/michaelrodriguess/auth_service/internal/service"
	"github.com/michaelrodriguess/auth_service/pkg/db/mongo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbClient := mongo.Connect(config.GetMongoURI())
	err = dbClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}
	db := dbClient.Database("authdb")

	repo := repository.NewUserAuthRepository(db)
	service := service.NewAuthService(repo)
	handler := handler.NewAuthHandler(service)

	r := gin.Default()
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	authMiddlewareGroup := r.Group("/")
	authMiddlewareGroup.Use(middleware.AuthMiddleware(repo))
	authMiddlewareGroup.GET("/me", handler.Me)
	authMiddlewareGroup.POST("/logout", handler.Logout)

	r.Run(":8080")
}
