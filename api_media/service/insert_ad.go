package service

import (
	"api_media/utils"
	"github.com/grafov/m3u8"
)

func (s mediaService) InsertAd(originalPlaylist *m3u8.MediaPlaylist, adPlaylist *m3u8.MediaPlaylist) {

	length := 0
	for _, segment := range originalPlaylist.Segments {
		if segment != nil {
			length += 1
		}
	}
	timeslots := utils.GetAdTimeSlots()

	for _, timeslot := range timeslots {
		insertPoint := int(float64(length) * timeslot)
		for i, segment := range adPlaylist.Segments {
			if segment != nil {
				originalPlaylist.Segments = append(originalPlaylist.Segments[:insertPoint+1], originalPlaylist.Segments[insertPoint:]...)
				originalPlaylist.Segments[insertPoint] = segment
				if i == 0 {
					originalPlaylist.Segments[insertPoint].Discontinuity = true
				}
				insertPoint++
			}
		}
		originalPlaylist.Segments[insertPoint].Discontinuity = true
	}
}
