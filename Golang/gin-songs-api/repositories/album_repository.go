package repositories

import (
	"context"
	"gin-songs-api/models/entity"
	"github.com/jackc/pgx/v4"
)

type AlbumRepository interface {
	Save(ctx context.Context, tx pgx.Tx, album *entity.Album) error
	FindById(ctx context.Context, tx pgx.Tx, id string) (*entity.Album, error)
	FindAll(ctx context.Context, tx pgx.Tx) ([]*entity.Album, error)
	Delete(ctx context.Context, tx pgx.Tx, id string) (int64, error)
}

type albumRepository struct{}

func (repo *albumRepository) Save(ctx context.Context, tx pgx.Tx, album *entity.Album) error {
	_, err := tx.Exec(ctx, `
		INSERT INTO albums (id, title, genre, artist_id, release_date)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE
		SET title = EXCLUDED.title,
			genre = EXCLUDED.genre,
			artist_id = EXCLUDED.artist_id,
			release_date = EXCLUDED.release_date,
			updated_at = NOW()
	`, album.ID, album.Title, album.Genre, album.ArtistID, album.ReleaseDate)

	return err
}

func (repo *albumRepository) FindById(ctx context.Context, tx pgx.Tx, id string) (*entity.Album, error) {
	row := tx.QueryRow(ctx, `
		SELECT id, title, genre, artist_id, release_date, created_at, updated_at
		FROM albums
		WHERE id = $1
	`, id)

	var album entity.Album
	err := row.Scan(&album.ID, &album.Title, &album.Genre, &album.ArtistID, &album.ReleaseDate, &album.CreatedAt, &album.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (repo *albumRepository) FindAll(ctx context.Context, tx pgx.Tx) ([]*entity.Album, error) {
	rows, err := tx.Query(ctx, `
		SELECT id, title, genre, artist_id, release_date, created_at, updated_at
		FROM albums
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []*entity.Album
	for rows.Next() {
		var album entity.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Genre, &album.ArtistID, &album.ReleaseDate, &album.CreatedAt, &album.UpdatedAt)
		if err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}

	return albums, nil
}

func (repo *albumRepository) Delete(ctx context.Context, tx pgx.Tx, id string) (int64, error) {
	commandTag, err := tx.Exec(ctx, `
		DELETE FROM albums
		WHERE id = $1
	`, id)
	if err != nil {
		return 0, err
	}

	return commandTag.RowsAffected(), nil
}

func NewAlbumRepository() AlbumRepository {
	return &albumRepository{}
}
