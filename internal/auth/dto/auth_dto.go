package dto

import (
	"github.com/theskinnycoder/streamlens/internal/repository"
)

type RegisterRequest struct {
	TenantName string `json:"tenant_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Message      string          `json:"message"`
	User         repository.User `json:"user"`
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
}

type AuthResponseWithoutRefreshToken struct {
	Message string          `json:"message"`
	User    repository.User `json:"user"`
	Token   string          `json:"token"`
}
