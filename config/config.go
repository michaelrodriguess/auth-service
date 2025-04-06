package config

import (
	"os"
)

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
