package service

import (
	"errors"
	"github.com/grafov/m3u8"
)

func (s mediaService) GetStreamMediaList(bucket string, file string) (*m3u8.MediaPlaylist, error) {
	object, err := s.m.Read(bucket, file)
	if err != nil {
		return &m3u8.MediaPlaylist{}, err
	}

	m3u8File, _, err := m3u8.DecodeFrom(object, true)
	if err != nil {
		return &m3u8.MediaPlaylist{}, err
	}
	playlist, ok := m3u8File.(*m3u8.MediaPlaylist)

	if !ok {
		return playlist, errors.New("error to decode file")
	}
	return playlist, nil
}
