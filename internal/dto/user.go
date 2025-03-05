package dto

import "github.com/google/uuid"

type UserCreateDTO struct {
	Username string
	Password string
}

type UserResponseDTO struct {
	Id uuid.UUID
	Username string
}