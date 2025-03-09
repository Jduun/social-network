package services

import (
	"github.com/google/uuid"
	"social-network/internal/dto"
)

type AuthService interface {
	CreateUser(user dto.UserCreateDTO) (dto.UserResponseDTO, error)
	LoginUser(userData dto.LoginUserDTO) (string, error)
	ParseToken(token string) (uuid.UUID, error)
}

type UserService interface {
	GetUserById(id uuid.UUID) (dto.UserResponseDTO, error)
	GetUsers() ([]dto.UserResponseDTO, error)
	GetFollowers() ([]dto.UserResponseDTO, error)
	GetFollowingUsers() ([]dto.UserResponseDTO, error)
}
