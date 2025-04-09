package service

import (
	"context"
	"errors"

	"github.com/michaelrodriguess/auth_service/internal/model"
	"github.com/michaelrodriguess/auth_service/internal/repository"
	"github.com/michaelrodriguess/auth_service/pkg/crypto"
	"github.com/michaelrodriguess/auth_service/pkg/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserAuthRepository
}

func NewAuthService(repo *repository.UserAuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req model.UserAuthRequest) (*model.AuthResponse, error) {
	hashedPassword, err := crypto.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.UserAuth{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
	}

	err = s.repo.CreateUserAuth(user)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token: token,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (s *AuthService) Login(req model.UserLoginRequest) (*model.LoginResponse, error) {
	ctx := context.TODO()

	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("email or password invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("email or password invalid")
	}

	token, err := jwt.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{Token: token}, nil
}

func (s *AuthService) GetProfile(userID string) (*model.UserAuth, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(context.TODO(), objID)
}
