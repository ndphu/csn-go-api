package service

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
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

func (*PlaylistService) FindAllPlaylist(page int, size int) ([]entity.Playlist, error) {
	return dao.GetAllPlaylist(page, size)
}
func (service *PlaylistService) CreatePlaylist(playlist *entity.Playlist) (error) {
	fmt.Println("creating playlist", playlist.Title, "with track count =", len(playlist.Tracks))
	playlist.Id = bson.NewObjectId()
	return dao.SavePlaylist(playlist)
}
func (service *PlaylistService) FindPlaylistById(id string) (*entity.Playlist, error) {
	return dao.FindPlaylistById(bson.ObjectIdHex(id))
}
