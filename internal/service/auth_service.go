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

func (s *AuthService) Logout(authHeader string) error {
	if authHeader == "" {
		return errors.New("authorization header is required")
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) {
		return errors.New("invalid authorization header format")
	}

	token := authHeader[len(bearerPrefix):]
	if token == "" {
		return errors.New("token not provided")
	}

	err := s.repo.AddTokenToBlocklist(context.TODO(), token)
	if err != nil {
		return errors.New("failed to invalidate token")
	}

	return nil
}

func (s *AuthService) ForgotPassword(req model.ForgotPasswordRequest) error {
	ctx := context.TODO()

	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return errors.New("email invalid")
	}

	err = crypto.ComparePassword(user.Password, req.OldPassword)
	if err != nil {
		return errors.New("old password invalid")
	}

	if req.OldPassword == req.NewPassword {
		return errors.New("new password cannot be the same as old password")
	}

	hashedPassword, err := crypto.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = s.repo.UpdateUserPassword(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
