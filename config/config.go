package config

// import (
// 	"os"
// )

func GetMongoURI() string {
	return "mongodb://localhost:27017"
	// return os.Getenv("MONGO_URI")
}
