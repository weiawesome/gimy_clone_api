package utils

const (
	MinIOUseSSL = false
	version     = "v1"

	playRoute   = "PlayRoute"
	fileID      = "FileID"
	fileEpisode = "FileEpisode"
	fileBucket  = "FileBucket"
	file        = "File"
	fileKey     = "FileKey"
)

var (
	adTimeSlots = []float64{0.3, 0.7}
)

func GetMinIOUseSSL() bool {
	return MinIOUseSSL
}
func GetAdTimeSlots() []float64 {
	return adTimeSlots
}

func GetVersion() string {
	return version
}
func GetPlayRouteParameter() string {
	return playRoute
}
func GetFileIDParameter() string {
	return fileID
}
func GetFileEpisodeParameter() string {
	return fileEpisode
}
func GetFileBucketParameter() string {
	return fileBucket
}
func GetFileParameter() string {
	return file
}
func GetFileKeyParameter() string {
	return fileKey
}
