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

type SongService interface {
	Create(ctx context.Context, song *entity.Song) (*api.Song, error)
	Update(ctx context.Context, song *entity.Song) (*api.Song, error)
	FindAll(ctx context.Context) ([]*api.Song, error)
	FindById(ctx context.Context, id string) (*api.Song, error)
	Delete(ctx context.Context, id string) error
}

type songService struct {
	songRepo repositories.SongRepository
	DB       *pgxpool.Pool
}

func (s *songService) Create(ctx context.Context, song *entity.Song) (*api.Song, error) {
	song.ID = utils.GenerateID()

	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	err = s.songRepo.Create(ctx, tx, song)
	if err != nil {
		return nil, err
	}

	data, err := s.songRepo.FindByID(ctx, tx, song.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPISong(data), nil
}

func (s *songService) Update(ctx context.Context, song *entity.Song) (*api.Song, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	err = s.songRepo.Update(ctx, tx, song)
	if err != nil {
		return nil, err
	}

	data, err := s.songRepo.FindByID(ctx, tx, song.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPISong(data), nil
}

func (s *songService) FindAll(ctx context.Context) ([]*api.Song, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	data, err := s.songRepo.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	var apiSongs []*api.Song
	for _, song := range data {
		apiSongs = append(apiSongs, utils.ConvertToAPISong(song))
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	if len(apiSongs) == 0 {
		return []*api.Song{}, nil
	}

	return apiSongs, nil
}

func (s *songService) FindById(ctx context.Context, id string) (*api.Song, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	data, err := s.songRepo.FindByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPISong(data), nil
}

func (s *songService) Delete(ctx context.Context, id string) error {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	rowsAffected, err := s.songRepo.Delete(ctx, tx, id)
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

func NewSongService(songRepo repositories.SongRepository, DB *pgxpool.Pool) SongService {
	return &songService{songRepo, DB}
}
