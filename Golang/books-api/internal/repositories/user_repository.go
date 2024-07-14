package repositories

import (
	"book-api/internal/models/entities"
	"context"
	"github.com/jackc/pgx/v4"
)

type UserRepository interface {
	Create(ctx context.Context, tx pgx.Tx, user *entities.User) error
	FindAll(ctx context.Context, tx pgx.Tx) ([]*entities.User, error)
	FindByID(ctx context.Context, tx pgx.Tx, id string) (*entities.User, error)
	FindByUsername(ctx context.Context, tx pgx.Tx, username string) (*entities.User, error)
}

type userRepository struct {
}

func (r *userRepository) Create(ctx context.Context, tx pgx.Tx, user *entities.User) error {
	query := `INSERT INTO users (full_name, username, password) VALUES ($1, $2, $3)`
	result, err := tx.Exec(ctx, query, user.FullName, user.Username, user.Password)
	if err != nil {
		return err
	}
	if result.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *userRepository) FindAll(ctx context.Context, tx pgx.Tx) ([]*entities.User, error) {
	query := `SELECT id, full_name, username, created_at, updated_at FROM users`
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := new(entities.User)
		err := rows.Scan(&user.ID, &user.FullName, &user.Username, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, tx pgx.Tx, id string) (*entities.User, error) {
	query := `SELECT id, full_name, username, password, created_at, updated_at FROM users WHERE id = $1`
	user := new(entities.User)
	err := tx.QueryRow(ctx, query, id).Scan(&user.ID, &user.FullName, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, tx pgx.Tx, username string) (*entities.User, error) {
	query := `SELECT id, full_name, username, password, created_at, updated_at FROM users WHERE username = $1`
	user := new(entities.User)
	err := tx.QueryRow(ctx, query, username).Scan(&user.ID, &user.FullName, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
