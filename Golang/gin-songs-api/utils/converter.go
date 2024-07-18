package utils

import (
	"gin-songs-api/models/api"
	"gin-songs-api/models/entity"
)

func ConvertToAPIArtist(artist *entity.Artist) *api.Artist {
	return &api.Artist{
		ID:        artist.ID,
		Name:      artist.Name,
		Bio:       artist.Bio,
		CreatedAt: artist.CreatedAt,
		UpdatedAt: artist.UpdatedAt,
	}
}

func ConvertToAPIArtists(artists []*entity.Artist) []*api.Artist {
	var apiArtists []*api.Artist
	for _, artist := range artists {
		apiArtists = append(apiArtists, ConvertToAPIArtist(artist))
	}
	return apiArtists
}

func ConvertToAPIAlbum(repoAlbum *entity.Album, repoArtist *entity.Artist) *api.Album {
	return &api.Album{
		ID:          repoAlbum.ID,
		Title:       repoAlbum.Title,
		Genre:       repoAlbum.Genre,
		Artist:      *ConvertToAPIArtist(repoArtist),
		ReleaseDate: repoAlbum.ReleaseDate,
		CreatedAt:   repoAlbum.CreatedAt,
		UpdatedAt:   repoAlbum.UpdatedAt,
	}
}

func ConvertToAPIAlbum2(repoAlbum *entity.Album) *api.Album {
	return &api.Album{
		ID:          repoAlbum.ID,
		Title:       repoAlbum.Title,
		Genre:       repoAlbum.Genre,
		Artist:      *ConvertToAPIArtist(&repoAlbum.Artist),
		ReleaseDate: repoAlbum.ReleaseDate,
		CreatedAt:   repoAlbum.CreatedAt,
		UpdatedAt:   repoAlbum.UpdatedAt,
	}
}

func ConvertToAPISong(repoSong *entity.Song) *api.Song {
	return &api.Song{
		ID:          repoSong.ID,
		Title:       repoSong.Title,
		Album:       *ConvertToAPIAlbum2(&repoSong.Album),
		Duration:    repoSong.Duration,
		ReleaseDate: repoSong.ReleaseDate,
		CreatedAt:   repoSong.CreatedAt,
		UpdatedAt:   repoSong.UpdatedAt,
	}
}
