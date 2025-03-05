package repositories

import (
	"context"
	"social-network/internal/dto"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user dto.UserCreateDTO) (dto.UserResponseDTO, error)
	GetUserById(ctx context.Context, id uuid.UUID) (dto.UserResponseDTO, error)
}
