package utils

import (
	"encoding/json"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"strconv"
)

type StreamInfo struct {
	Streams []Stream `json:"streams"`
}
type Stream struct {
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	BitRate string `json:"bit_rate"`
}

func (s Stream) GetBitRate() (int, error) {
	return strconv.Atoi(s.BitRate)
}

func GetVideoStream(filePath string, fileName string) (Stream, error) {
	data, err := ffmpeg.Probe(filePath + "/" + fileName)
	if err != nil {
		return Stream{}, err
	}
	var info StreamInfo
	err = json.Unmarshal([]byte(data), &info)

	return info.Streams[0], err
}

func TranslateToHLS(filePath string, fileName string, translateName string, stream Stream) error {
	return ffmpeg.Input(filePath+"/"+fileName).
		Output(filePath+"/"+translateName+".m3u8", ffmpeg.KwArgs{
			"vprofile":             "baseline",
			"level":                "3.0",
			"s":                    strconv.Itoa(stream.Width) + "x" + strconv.Itoa(stream.Height),
			"start_number":         0,
			"hls_time":             10,
			"hls_list_size":        0,
			"f":                    "hls",
			"hls_segment_filename": filePath + "/resource/" + translateName + "-" + "%d.ts",
		}).OverWriteOutput().Run()
}
