package service

import (
	"github.com/michaelrodriguess/auth_service/internal/model"
	"github.com/michaelrodriguess/auth_service/internal/repository"
	"github.com/michaelrodriguess/auth_service/pkg/crypto"
	"github.com/michaelrodriguess/auth_service/pkg/jwt"
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

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token: token,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
