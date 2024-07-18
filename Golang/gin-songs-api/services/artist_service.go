package services

import (
	"context"
	"gin-songs-api/models/api"
	"gin-songs-api/models/entity"
	"gin-songs-api/repositories"
	"gin-songs-api/utils"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ArtistService interface {
	Create(ctx context.Context, artist *entity.Artist) (*api.Artist, error)
	Update(ctx context.Context, id string, artist *entity.Artist) (*api.Artist, error)
	FindAll(ctx context.Context) ([]*api.Artist, error)
	FindById(ctx context.Context, id string) (*api.Artist, error)
	Delete(ctx context.Context, id string) error
}

type artistService struct {
	artistRepo repositories.ArtistRepository
	DB         *pgxpool.Pool
}

func (s *artistService) Create(ctx context.Context, artist *entity.Artist) (*api.Artist, error) {
	artist.ID = utils.GenerateID()

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			// TODO: return error
			return
		}
	}(tx, ctx)

	err = s.artistRepo.Save(ctx, tx, artist)
	if err != nil {
		return nil, err
	}

	data, err := s.artistRepo.FindById(ctx, tx, artist.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPIArtist(data), nil
}

func (s *artistService) Update(ctx context.Context, id string, artist *entity.Artist) (*api.Artist, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			// TODO: return error
			return
		}
	}(tx, ctx)

	artist.ID = id

	_, err = s.artistRepo.FindById(ctx, tx, artist.ID)
	if err != nil {
		return nil, err
	}

	err = s.artistRepo.Save(ctx, tx, artist)
	if err != nil {
		return nil, err
	}

	data, err := s.artistRepo.FindById(ctx, tx, artist.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPIArtist(data), nil
}

func (s *artistService) FindAll(ctx context.Context) ([]*api.Artist, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			// TODO: return error
			return
		}
	}(tx, ctx)

	data, err := s.artistRepo.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPIArtists(data), nil
}

func (s *artistService) FindById(ctx context.Context, id string) (*api.Artist, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			// TODO: return error
			return
		}
	}(tx, ctx)

	data, err := s.artistRepo.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPIArtist(data), nil
}

func (s *artistService) Delete(ctx context.Context, id string) error {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			// TODO: return error
			return
		}
	}(tx, ctx)

	rowsAffected, err := s.artistRepo.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return pgx.ErrNoRows
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func NewArtistService(artistRepo repositories.ArtistRepository, DB *pgxpool.Pool) ArtistService {
	return &artistService{
		artistRepo: artistRepo,
		DB:         DB,
	}
}
