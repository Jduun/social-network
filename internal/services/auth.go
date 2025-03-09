package services

import (
	"context"
	"errors"
	"time"

	"social-network/config"
	"social-network/internal/dto"
	"social-network/internal/repositories"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthServiceImpl(userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{userRepo: userRepo}
}

func (s *AuthServiceImpl) CreateUser(user dto.UserCreateDTO) (dto.UserResponseDTO, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponseDTO{}, ErrPasswordHashing
	}
	user.Password = string(hashedBytes)

	created_user, err := s.userRepo.CreateUser(context.Background(), user)
	if err != nil {
		if errors.Is(err, repositories.ErrAlreadyExists) {
			return dto.UserResponseDTO{}, ErrUserAlreadyExists
		} else if errors.Is(err, repositories.ErrNotFound) {
			return dto.UserResponseDTO{}, ErrUserNotFound
		} else {
			return dto.UserResponseDTO{}, ErrCannotCreateUser
		}
	}
	return created_user, nil
}

func (s *AuthServiceImpl) LoginUser(userData dto.LoginUserDTO) (string, error) {
	user, err := s.userRepo.GetUserByName(context.Background(), userData.Username)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return "", ErrCannotFindUser
		} else {
			return "", ErrCannotLoginUser
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.Cfg.JwtSecret))
	if err != nil {
		return "", ErrCannotSignToken
	}
	return tokenString, nil
}

func (s *AuthServiceImpl) ParseToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JwtSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		userId, err := uuid.Parse((*claims)["user_id"].(string))
		if err != nil {
			return uuid.UUID{}, ErrInvalidToken
		}
		return userId, nil
	} else {
		return uuid.UUID{}, ErrInvalidToken
	}
}
