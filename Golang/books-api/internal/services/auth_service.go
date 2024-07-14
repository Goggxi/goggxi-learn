package services

import (
	"book-api/internal/models/entities"
	"book-api/internal/models/responses"
	"book-api/internal/repositories"
	"book-api/pkg/utils"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthService interface {
	Signup(ctx context.Context, user *entities.User) (*responses.UserTokenResponse, error)
	Login(ctx context.Context, username, password string) (*responses.UserTokenResponse, error)
	Logout(ctx context.Context, userID string) error
	GetCurrentUser(ctx context.Context, userID string) (*responses.UserResponse, error)
	RefreshToken(token string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
	db       *pgxpool.Pool
}

func (s *authService) Signup(ctx context.Context, user *entities.User) (*responses.UserTokenResponse, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	err = s.userRepo.Create(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	userRes, err := s.userRepo.FindByUsername(ctx, tx, user.Username)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(userRes.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &responses.UserTokenResponse{
		User: responses.UserResponse{
			ID:       userRes.ID,
			FullName: userRes.FullName,
			Username: userRes.Username,
		},
		Token: token,
	}, nil
}

func (s *authService) Login(ctx context.Context, username, password string) (*responses.UserTokenResponse, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	user, err := s.userRepo.FindByUsername(ctx, tx, username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &responses.UserTokenResponse{
		User: responses.UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Username: user.Username,
		},
		Token: token,
	}, nil
}

func (s *authService) Logout(ctx context.Context, userID string) error {
	// Implement token invalidation if needed
	return nil
}

func (s *authService) GetCurrentUser(ctx context.Context, userID string) (*responses.UserResponse, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	user, err := s.userRepo.FindByID(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &responses.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
	}, nil
}

func (s *authService) RefreshToken(token string) (string, error) {
	refreshToken, err := utils.RefreshToken(token)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func NewAuthService(userRepo repositories.UserRepository, db *pgxpool.Pool) AuthService {
	return &authService{
		userRepo: userRepo,
		db:       db,
	}
}
