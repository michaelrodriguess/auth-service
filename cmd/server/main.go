package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/michaelrodriguess/auth_service/config"
	_ "github.com/michaelrodriguess/auth_service/docs"
	"github.com/michaelrodriguess/auth_service/internal/handler"
	"github.com/michaelrodriguess/auth_service/internal/middleware"
	"github.com/michaelrodriguess/auth_service/internal/repository"
	"github.com/michaelrodriguess/auth_service/internal/service"
	"github.com/michaelrodriguess/auth_service/pkg/db/mongo"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// @title Auth Service API
// @version 1.0
// @description API para autenticação de usuários
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/forgot-password", handler.ForgotPassword)

	authMiddlewareGroup := r.Group("/")
	authMiddlewareGroup.Use(middleware.AuthMiddleware(repo))
	authMiddlewareGroup.GET("/me", handler.Me)
	authMiddlewareGroup.POST("/logout", handler.Logout)

	log.Println("✅ Auth service running on port 8080")
	log.Println("📖 Swagger docs available at: http://localhost:8080/swagger/index.html")

	r.Run(":8080")
}
