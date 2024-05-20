package service

import (
	"github.com/grafov/m3u8"
)

func (s mediaService) RefreshBucketUrl(playlist *m3u8.MediaPlaylist, bucket string) {
	if playlist.Key != nil {
		playlist.Key.URI = bucket + "/" + playlist.Key.URI
	}
	for _, segment := range playlist.Segments {
		if segment != nil {
			segment.URI = bucket + "/" + segment.URI
		}
	}
}
