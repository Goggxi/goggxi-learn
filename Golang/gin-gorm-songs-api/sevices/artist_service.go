package sevices

import (
	"gin-gorm-songs-api/models/apis"
	"gin-gorm-songs-api/models/entities"
	"gin-gorm-songs-api/repositories"
	"gin-gorm-songs-api/utils"
	"gorm.io/gorm"
)

type ArtistService interface {
	Create(artist *apis.ArtistReq) (*apis.ArtistRes, error)
	Update(artist *apis.ArtistReq) (*apis.ArtistRes, error)
	Delete(id int) error
	GetById(id int) (*apis.ArtistRes, error)
	GetAll() ([]apis.ArtistRes, error)
}

type artistService struct {
	artistRepo repositories.ArtisRepository
	db         *gorm.DB
}

func (s *artistService) Create(artist *apis.ArtistReq) (*apis.ArtistRes, error) {
	id := utils.GenerateID()

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	artistEntity := entities.Artist{
		ID:   id,
		Name: artist.Name,
		Bio:  artist.Bio,
	}

	if err := s.artistRepo.Save(tx, &artistEntity); err != nil {
		tx.Rollback()
		return nil, err
	}

	byId, err := s.artistRepo.FindById(tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &apis.ArtistRes{
		ID:        byId.ID,
		Name:      byId.Name,
		Bio:       byId.Bio,
		CreatedAt: byId.CreatedAt,
		UpdatedAt: byId.UpdatedAt,
	}, nil
}

func (s *artistService) Update(artist *apis.ArtistReq) (*apis.ArtistRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *artistService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *artistService) GetById(id int) (*apis.ArtistRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *artistService) GetAll() ([]apis.ArtistRes, error) {
	//TODO implement me
	panic("implement me")
}

func NewArtistService(artistRepo repositories.ArtisRepository, db *gorm.DB) ArtistService {
	return &artistService{
		artistRepo: artistRepo,
		db:         db,
	}
}
