package hashing

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type HashingService struct {
	hashCost int
}

func NewHashingService(hashCost int) *HashingService {
	return &HashingService{hashCost: hashCost}
}

func (s *HashingService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.hashCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(hashedPassword), nil
}

func (s *HashingService) ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
