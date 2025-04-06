package main

import (
	"context"
	"log"

	"github.com/michaelrodriguess/auth_service/config"
	"github.com/michaelrodriguess/auth_service/internal/handler"
	"github.com/michaelrodriguess/auth_service/internal/repository"
	"github.com/michaelrodriguess/auth_service/internal/service"
	"github.com/michaelrodriguess/auth_service/pkg/db/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	dbClient := mongo.Connect(config.GetMongoURI())
	err := dbClient.Ping(context.TODO(), readpref.Primary())
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

	r.Run(":8080")
}
