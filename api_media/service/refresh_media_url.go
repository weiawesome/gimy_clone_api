package service

import (
	"api_media/utils"
	"github.com/grafov/m3u8"
)

func (s mediaService) RefreshMediaUrl(playlist *m3u8.MediaPlaylist) {
	for _, segment := range playlist.Segments {
		if segment != nil {
			segment.URI = utils.EnvCDNAddress() + "/api/" + utils.GetVersion() + "/resource/media/" + segment.URI
		}
	}
}
