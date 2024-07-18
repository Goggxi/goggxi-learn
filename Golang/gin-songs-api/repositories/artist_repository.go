package repositories

import (
	"context"
	"gin-songs-api/models/entity"
	"github.com/jackc/pgx/v4"
)

type ArtistRepository interface {
	Save(ctx context.Context, tx pgx.Tx, artist *entity.Artist) error
	FindAll(ctx context.Context, tx pgx.Tx) ([]*entity.Artist, error)
	FindById(ctx context.Context, tx pgx.Tx, id string) (*entity.Artist, error)
	Delete(ctx context.Context, tx pgx.Tx, id string) (int64, error)
}

type artistRepository struct {
}

func (r *artistRepository) Save(ctx context.Context, tx pgx.Tx, artist *entity.Artist) error {
	sql := `INSERT INTO artists (id, name, bio) VALUES ($1, $2, $3) 
			ON CONFLICT (id) 
			DO UPDATE SET name = $2, bio = $3, updated_at = NOW()`

	_, err := tx.Exec(ctx, sql, artist.ID, artist.Name, artist.Bio)

	return err
}

func (r *artistRepository) FindAll(ctx context.Context, tx pgx.Tx) ([]*entity.Artist, error) {
	sql := `SELECT id, name, bio, created_at, updated_at FROM artists`
	rows, err := tx.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	var artists []*entity.Artist
	for rows.Next() {
		artist := new(entity.Artist)
		err := rows.Scan(&artist.ID, &artist.Name, &artist.Bio, &artist.CreatedAt, &artist.UpdatedAt)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

func (r *artistRepository) FindById(ctx context.Context, tx pgx.Tx, id string) (*entity.Artist, error) {
	sql := `SELECT id, name, bio, created_at, updated_at FROM artists WHERE id = $1`
	artist := new(entity.Artist)

	err := tx.QueryRow(ctx, sql, id).Scan(&artist.ID, &artist.Name, &artist.Bio, &artist.CreatedAt, &artist.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (r *artistRepository) Delete(ctx context.Context, tx pgx.Tx, id string) (int64, error) {
	sql := `DELETE FROM artists WHERE id = $1`
	commandTag, err := tx.Exec(ctx, sql, id)
	if err != nil {
		return 0, err
	}

	return commandTag.RowsAffected(), nil
}

func NewArtistRepository() ArtistRepository {
	return &artistRepository{}
}
