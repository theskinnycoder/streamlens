package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/theskinnycoder/streamlens/internal/auth/dto"
	"github.com/theskinnycoder/streamlens/internal/hashing"
	"github.com/theskinnycoder/streamlens/internal/jwt"
	"github.com/theskinnycoder/streamlens/internal/repository"
)

type AuthService struct {
	repo           *repository.Queries
	hashingService *hashing.HashingService
	jwtService     *jwt.JWTService
}

func NewAuthService(repo *repository.Queries, hashing *hashing.HashingService, jwt *jwt.JWTService) *AuthService {
	return &AuthService{repo: repo, hashingService: hashing, jwtService: jwt}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (dto.AuthResponse, error) {
	// 1. Check if user already exists
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return dto.AuthResponse{}, err
	}
	if user.ID != uuid.Nil {
		return dto.AuthResponse{}, errors.New("user already exists")
	}

	// 2. Hash password
	hashedPassword, err := s.hashingService.HashPassword(req.Password)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 3. Check if tenant exists
	tenant, err := s.repo.GetTenantByName(ctx, req.TenantName)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return dto.AuthResponse{}, err
	}

	// 4. If tenant does not exist, create it
	if tenant.ID == uuid.Nil {
		tenant, err = s.repo.CreateTenant(ctx, req.TenantName)
		if err != nil {
			return dto.AuthResponse{}, err
		}
	}

	// 5. Create user
	createdUser, err := s.repo.CreateUser(ctx, repository.CreateUserParams{
		TenantID:       tenant.ID,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 6. Create user role
	_, err = s.repo.AssignUserRole(ctx, repository.AssignUserRoleParams{
		UserID: createdUser.ID,
		Role:   "admin",
	})
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 7. Generate access token
	accessToken, err := s.jwtService.GenerateToken(createdUser.ID)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 8. Generate refresh token
	refreshToken, err := s.jwtService.GenerateToken(createdUser.ID)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 9. Return response
	return dto.AuthResponse{
		Message:      "user created successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: repository.User{
			ID:        createdUser.ID,
			TenantID:  createdUser.TenantID,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		},
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error) {
	// 1. Check if user exists
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 2. Check if password is correct
	err = s.hashingService.ComparePassword(user.HashedPassword, req.Password)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 3. Generate access token
	accessToken, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 4. Generate refresh token
	refreshToken, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// 5. Return response
	return dto.AuthResponse{
		Message:      "user logged in successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: repository.User{
			ID:        user.ID,
			TenantID:  user.TenantID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}
