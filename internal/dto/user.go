package dto

import "github.com/google/uuid"

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateDTO struct {
	LoginUserDTO
}

type UserResponseDTO struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type UserEntityDTO struct {
	UserResponseDTO
	Password string `json:"password"`
}
