package repositories

import (
	"context"
	"social-network/internal/dto"
	"social-network/pkg/database"

	"github.com/google/uuid"
)

type UserPostgresRepository struct {
	db *database.PostgresDatabase
}

func NewUserPostgresRepository(db *database.PostgresDatabase) UserRepository {
	return &UserPostgresRepository{db: db}
}

func (repo *UserPostgresRepository) CreateUser(ctx context.Context, user dto.UserCreateDTO) (dto.UserResponseDTO, error) {
	query := `
		INSERT INTO "user" (username, password) 
		VALUES ($1, $2)
		RETURNING id, username`

	var createdUser dto.UserResponseDTO
	err := repo.db.Pool.
		QueryRow(ctx, query, user.Username, user.Password).
		Scan(&createdUser.Id, &createdUser.Username)

	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return createdUser, nil
}

func (repo *UserPostgresRepository) GetUserById(ctx context.Context, id uuid.UUID) (dto.UserResponseDTO, error) {
	query := `
		SELECT id, username
		FROM "user"
		WHERE id = $1`

	var user dto.UserResponseDTO
	err := repo.db.Pool.
		QueryRow(ctx, query, id).
		Scan(&user.Id, &user.Username)

	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return user, nil
}
