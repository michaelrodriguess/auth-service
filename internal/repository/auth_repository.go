package repository

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/michaelrodriguess/auth_service/config"
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

func (r *UserAuthRepository) CreateUserAuth(user *model.UserAuth) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	_, err := r.GetByEmail(context.TODO(), user.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	_, err = r.collection.InsertOne(context.TODO(), user)
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

func (r *UserAuthRepository) AddTokenToBlocklist(ctx context.Context, token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTSecret()), nil
	})

	if err != nil {
		return err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	var expiresAt time.Time
	if exp, ok := claims["exp"].(float64); ok {
		expiresAt = time.Unix(int64(exp), 0)
	} else {
		expiresAt = time.Now().Add(time.Hour * 24)
	}

	blockedTokensColl := r.collection.Database().Collection("blocked_tokens")

	_, err = blockedTokensColl.InsertOne(ctx, model.BlockedToken{
		Token:     token,
		ExpiresAt: expiresAt,
	})

	return err
}

func (r *UserAuthRepository) IsTokenBlocked(ctx context.Context, token string) (bool, error) {
	blockedTokensColl := r.collection.Database().Collection("blocked_tokens")

	count, err := blockedTokensColl.CountDocuments(ctx, bson.M{"token": token})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserAuthRepository) UpdateUserPassword(ctx context.Context, user *model.UserAuth) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user, "$currentDate": bson.M{"updatedAt": true}})
	if err != nil {
		return err
	}

	return nil
}
