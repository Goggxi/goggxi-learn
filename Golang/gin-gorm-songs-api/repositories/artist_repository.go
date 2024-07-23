package repositories

import (
	"gin-gorm-songs-api/models/entities"
	"gorm.io/gorm"
)

type ArtisRepository interface {
	Save(tx *gorm.DB, artist *entities.Artist) error
	FindById(tx *gorm.DB, id string) (*entities.Artist, error)
	FindAll(tx *gorm.DB) ([]entities.Artist, error)
	Delete(tx *gorm.DB, artist *entities.Artist) error
}

type artistRepository struct{}

func (r *artistRepository) Save(tx *gorm.DB, artist *entities.Artist) error {
	if err := tx.Create(artist).Error; err != nil {
		return err
	}

	return nil
}

func (r *artistRepository) FindById(tx *gorm.DB, id string) (*entities.Artist, error) {
	var artist entities.Artist
	if err := tx.First(&artist, id).Error; err != nil {
		return nil, err
	}

	return &artist, nil
}

func (r *artistRepository) FindAll(tx *gorm.DB) ([]entities.Artist, error) {
	var artists []entities.Artist
	if err := tx.Find(&artists).Error; err != nil {
		return nil, err
	}

	return artists, nil
}

func (r *artistRepository) Delete(tx *gorm.DB, artist *entities.Artist) error {
	if err := tx.Delete(artist).Error; err != nil {
		return err
	}

	return nil
}

func NewArtisRepository() ArtisRepository {
	return &artistRepository{}
}
