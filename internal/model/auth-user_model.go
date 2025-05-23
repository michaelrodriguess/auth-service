package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuth struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `json:"email"`
	Password  string             `json:"-" bson:"password"`
	Role      string             `json:"role" bson:"role"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

type UserAuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

type AuthResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type BlockedToken struct {
	Token     string    `bson:"token"`
	ExpiresAt time.Time `bson:"expires_at"`
}

type ForgotPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	OldPassword string `json:"old_password" binding:"required,min=8"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}
