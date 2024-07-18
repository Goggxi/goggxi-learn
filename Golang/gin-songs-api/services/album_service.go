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

type AlbumService interface {
	Create(ctx context.Context, album *entity.Album) (*api.Album, error)
	Update(ctx context.Context, album *entity.Album) (*api.Album, error)
	FindAll(ctx context.Context) ([]*api.Album, error)
	FindById(ctx context.Context, id string) (*api.Album, error)
	Delete(ctx context.Context, id string) error
}

type albumService struct {
	albumRepo  repositories.AlbumRepository
	artistRepo repositories.ArtistRepository
	DB         *pgxpool.Pool
}

func (s *albumService) Create(ctx context.Context, album *entity.Album) (*api.Album, error) {
	album.ID = utils.GenerateID()

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

	err = s.albumRepo.Save(ctx, tx, album)
	if err != nil {
		return nil, err
	}

	data, err := s.albumRepo.FindById(ctx, tx, album.ID)
	if err != nil {
		return nil, err
	}

	artist, err := s.artistRepo.FindById(ctx, tx, album.Artist.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPIAlbum(data, artist), nil
}

func (s *albumService) Update(ctx context.Context, album *entity.Album) (*api.Album, error) {
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

	_, err = s.albumRepo.FindById(ctx, tx, album.ID)
	if err != nil {
		return nil, err
	}

	err = s.albumRepo.Save(ctx, tx, album)
	if err != nil {
		return nil, err
	}

	data, err := s.albumRepo.FindById(ctx, tx, album.ID)
	if err != nil {
		return nil, err
	}

	artist, err := s.artistRepo.FindById(ctx, tx, album.Artist.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPIAlbum(data, artist), nil
}

func (s *albumService) FindAll(ctx context.Context) ([]*api.Album, error) {
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

	data, err := s.albumRepo.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	var apiAlbums []*api.Album
	for _, album := range data {
		artist, err := s.artistRepo.FindById(ctx, tx, album.Artist.ID)
		if err != nil {
			return nil, err
		}
		apiAlbums = append(apiAlbums, utils.ConvertToAPIAlbum(album, artist))
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	if len(apiAlbums) == 0 {
		return []*api.Album{}, nil
	}

	return apiAlbums, nil
}

func (s *albumService) FindById(ctx context.Context, id string) (*api.Album, error) {
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

	data, err := s.albumRepo.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	artist, err := s.artistRepo.FindById(ctx, tx, data.Artist.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertToAPIAlbum(data, artist), nil
}

func (s *albumService) Delete(ctx context.Context, id string) error {
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

	rowsAffected, err := s.albumRepo.Delete(ctx, tx, id)
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

func NewAlbumService(albumRepo repositories.AlbumRepository, artistRepo repositories.ArtistRepository, DB *pgxpool.Pool) AlbumService {
	return &albumService{albumRepo, artistRepo, DB}
}
