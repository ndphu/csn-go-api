package service

import (
	"github.com/ndphu/csn-go-api/dao"
	"github.com/ndphu/csn-go-api/entity"
)

type PlaylistService struct {
}

var playlistService *PlaylistService

func GetPlaylistService() *PlaylistService  {
	if playlistService == nil {
		playlistService = &PlaylistService{}
	}

	return playlistService
}

func (*PlaylistService) FindAllPlaylist(page int) ([]entity.Playlist, error) {
	return dao.GetAllPlaylist(page)
}
