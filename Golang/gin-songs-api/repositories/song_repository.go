package repositories

import (
	"context"
	"gin-songs-api/models/entity"
	"github.com/jackc/pgx/v4"
)

type SongRepository interface {
	Create(ctx context.Context, tx pgx.Tx, song *entity.Song) error
	Update(ctx context.Context, tx pgx.Tx, song *entity.Song) error
	FindAll(ctx context.Context, tx pgx.Tx) ([]*entity.Song, error)
	FindByID(ctx context.Context, tx pgx.Tx, id string) (*entity.Song, error)
	Delete(ctx context.Context, tx pgx.Tx, id string) (int64, error)
}

type songRepository struct{}

func (repo *songRepository) Create(ctx context.Context, tx pgx.Tx, song *entity.Song) error {
	_, err := tx.Exec(ctx, `
		INSERT INTO songs (id, title, album_id, duration, release_date)
		VALUES ($1, $2, $3, $4, $5)
	`, song.ID, song.Title, song.Album.ID, song.Duration, song.ReleaseDate)

	return err
}

func (repo *songRepository) Update(ctx context.Context, tx pgx.Tx, song *entity.Song) error {
	_, err := tx.Exec(ctx, `
		UPDATE songs
		SET title = $1, album_id = $2, duration = $3, release_date = $4, updated_at = NOW()
		WHERE id = $5
	`, song.Title, song.Album.ID, song.Duration, song.ReleaseDate, song.ID)

	return err
}

func (repo *songRepository) FindAll(ctx context.Context, tx pgx.Tx) ([]*entity.Song, error) {
	rows, err := tx.Query(ctx, `
		SELECT s.id, s.title, s.album_id, s.duration, s.release_date, s.created_at, s.updated_at,
			   a.id as album_id, a.title as album_title, a.genre as album_genre, a.artist_id as album_artist_id, a.release_date as album_release_date, a.created_at as album_created_at, a.updated_at as album_updated_at,
			   ar.id as artist_id, ar.name as artist_name, ar.bio as artist_bio, ar.created_at as artist_created_at, ar.updated_at as artist_updated_at
		FROM songs s
		JOIN albums a ON s.album_id = a.id
		JOIN artists ar ON a.artist_id = ar.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*entity.Song
	for rows.Next() {
		var song entity.Song

		err := rows.Scan(
			&song.ID, &song.Title, &song.Album.ID, &song.Duration, &song.ReleaseDate, &song.CreatedAt, &song.UpdatedAt,
			&song.Album.ID, &song.Album.Title, &song.Album.Genre, &song.Album.Artist.ID, &song.Album.ReleaseDate, &song.Album.CreatedAt, &song.Album.UpdatedAt,
			&song.Album.Artist.ID, &song.Album.Artist.Name, &song.Album.Artist.Bio, &song.Album.Artist.CreatedAt, &song.Album.Artist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		songs = append(songs, &song)
	}

	return songs, nil
}

func (repo *songRepository) FindByID(ctx context.Context, tx pgx.Tx, id string) (*entity.Song, error) {
	row := tx.QueryRow(ctx, `
		SELECT s.id, s.title, s.album_id, s.duration, s.release_date, s.created_at, s.updated_at,
			   a.id as album_id, a.title as album_title, a.genre as album_genre, a.artist_id as album_artist_id, a.release_date as album_release_date, a.created_at as album_created_at, a.updated_at as album_updated_at,
			   ar.id as artist_id, ar.name as artist_name, ar.bio as artist_bio, ar.created_at as artist_created_at, ar.updated_at as artist_updated_at
		FROM songs s
		JOIN albums a ON s.album_id = a.id
		JOIN artists ar ON a.artist_id = ar.id
		WHERE s.id = $1
	`, id)

	var song entity.Song
	err := row.Scan(
		&song.ID, &song.Title, &song.Album.ID, &song.Duration, &song.ReleaseDate, &song.CreatedAt, &song.UpdatedAt,
		&song.Album.ID, &song.Album.Title, &song.Album.Genre, &song.Album.Artist.ID, &song.Album.ReleaseDate, &song.Album.CreatedAt, &song.Album.UpdatedAt,
		&song.Album.Artist.ID, &song.Album.Artist.Name, &song.Album.Artist.Bio, &song.Album.Artist.CreatedAt, &song.Album.Artist.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (repo *songRepository) Delete(ctx context.Context, tx pgx.Tx, id string) (int64, error) {
	commandTag, err := tx.Exec(ctx, `
		DELETE FROM songs
		WHERE id = $1
	`, id)
	if err != nil {
		return 0, err
	}

	return commandTag.RowsAffected(), nil
}

func NewSongRepository() SongRepository {
	return &songRepository{}
}
