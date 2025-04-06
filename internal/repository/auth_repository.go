package repository

import (
	"context"
	"errors"
	"time"

	"github.com/michaelrodriguess/auth_service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
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
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserAuthRepository) GetByEmail(ctx context.Context, email string) (*model.UserAuth, error) {
	var user model.UserAuth

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("users not found")
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserAuthRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.UserAuth, error) {
	var user model.UserAuth

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
