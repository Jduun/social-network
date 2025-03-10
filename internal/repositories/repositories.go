package repositories

import (
	"context"
	"github.com/google/uuid"
	"social-network/internal/dto"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user dto.UserCreateDTO) (dto.UserResponseDTO, error)
	GetUserById(ctx context.Context, id uuid.UUID) (dto.UserEntityDTO, error)
	GetUserByName(ctx context.Context, username string) (dto.UserEntityDTO, error)
}
