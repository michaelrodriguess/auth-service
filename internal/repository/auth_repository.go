package repository

import (
	"context"
	"time"

	"github.com/michaelrodriguess/auth_service/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserAuthRepository struct {
	collection *mongo.Collection
}

func NewUserAuthRepository(db *mongo.Database) *UserAuthRepository {
	return &UserAuthRepository{
		collection: db.Collection("users_auth"),
	}
}

func (r *UserAuthRepository) Create(user *model.UserAuth) error {
	user.ID = primitive.NewObjectID().Hex()
	user.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}
