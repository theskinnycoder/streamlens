package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	secretKey string
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{secretKey: secretKey}
}

func (s *JWTService) GenerateToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
}

func (s *JWTService) GetUserID(token *jwt.Token) (uuid.UUID, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return uuid.UUID{}, errors.New("user ID is missing")
	}

	return uuid.MustParse(userID), nil
}

func (s *JWTService) GetUserRoles(token *jwt.Token) ([]string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return []string{}, errors.New("invalid token claims")
	}

	roles, ok := claims["roles"].([]interface{})
	if !ok {
		return []string{}, errors.New("roles are missing")
	}

	var roleStrings []string
	for _, role := range roles {
		roleString, ok := role.(string)
		if !ok {
			return []string{}, errors.New("invalid role type")
		}
		roleStrings = append(roleStrings, roleString)
	}

	return roleStrings, nil
}
