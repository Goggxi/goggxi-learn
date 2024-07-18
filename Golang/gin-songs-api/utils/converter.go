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

func ConvertToAPIAlbums(repoAlbums []*entity.Album, repoArtists []*entity.Artist) []*api.Album {
	var apiAlbums []*api.Album
	for i, repoAlbum := range repoAlbums {
		apiAlbums = append(apiAlbums, ConvertToAPIAlbum(repoAlbum, repoArtists[i]))
	}
	return apiAlbums
}

func ConvertToAPISong(repoSong *entity.Song, repoAlbum *entity.Album, repoArtist *entity.Artist) *api.Song {
	return &api.Song{
		ID:          repoSong.ID,
		Title:       repoSong.Title,
		Album:       *ConvertToAPIAlbum(repoAlbum, repoArtist),
		Duration:    repoSong.Duration,
		ReleaseDate: repoSong.ReleaseDate,
		CreatedAt:   repoSong.CreatedAt,
		UpdatedAt:   repoSong.UpdatedAt,
	}
}

func ConvertToAPISongs(repoSongs []*entity.Song, repoAlbums []*entity.Album, repoArtists []*entity.Artist) []*api.Song {
	var apiSongs []*api.Song
	for i, repoSong := range repoSongs {
		apiSongs = append(apiSongs, ConvertToAPISong(repoSong, repoAlbums[i], repoArtists[i]))
	}
	return apiSongs
}
