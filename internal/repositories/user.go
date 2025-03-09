package repositories

import (
	"context"
	"errors"
	"fmt"

	"social-network/internal/dto"
	"social-network/pkg/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserPostgresRepository struct {
	db *database.PostgresDatabase
}

func NewUserPostgresRepository(db *database.PostgresDatabase) UserRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) CreateUser(ctx context.Context, user dto.UserCreateDTO) (dto.UserResponseDTO, error) {
	sql := `
		insert into "user" (username, password) 
		values ($1, $2)
		returning id, username`

	var createdUser dto.UserResponseDTO
	err := r.db.Pool.
		QueryRow(ctx, sql, user.Username, user.Password).
		Scan(&createdUser.Id, &createdUser.Username)

	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return dto.UserResponseDTO{}, ErrAlreadyExists
			}
		}
		return dto.UserResponseDTO{}, fmt.Errorf("UserRepo.CreateUser error: %v", err)
	}
	return createdUser, nil
}

func (r *UserPostgresRepository) GetUserById(ctx context.Context, id uuid.UUID) (dto.UserEntityDTO, error) {
	sql := `
		select id, username
		from "user"
		where id = $1`

	var user dto.UserEntityDTO
	err := r.db.Pool.
		QueryRow(ctx, sql, id).
		Scan(&user.Id, &user.Username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.UserEntityDTO{}, ErrNotFound
		}
		return dto.UserEntityDTO{}, fmt.Errorf("UserRepo.GetUserById error: %v", err)
	}
	return user, nil
}

func (r *UserPostgresRepository) GetUserByName(ctx context.Context, username string) (dto.UserEntityDTO, error) {
	sql := `
		select id, username, password
		from "user"
		where username = $1`

	var user dto.UserEntityDTO
	err := r.db.Pool.
		QueryRow(ctx, sql, username).
		Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.UserEntityDTO{}, ErrNotFound
		}
		return dto.UserEntityDTO{}, fmt.Errorf("UserRepo.GetUserByName error: %v", err)
	}
	return user, nil
}
