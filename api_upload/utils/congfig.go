package utils

const (
	MinIOUseSSL = false
	version     = "v1"

	playRoute   = "PlayRoute"
	fileID      = "FileID"
	fileEpisode = "FileEpisode"
	state       = "state"

	file       = "file"
	adType     = "advertisement_type"
	expireTime = "expired_time"

	adBucket            = "advertisement"
	imageResourceBucket = "image"

	originalFileDescription = "-original-file"

	defaultValue = ""
)

func GetMinIOUseSSL() bool {
	return MinIOUseSSL
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
func GetStateParameter() string {
	return state
}

func GetFileParameter() string {
	return file
}
func GetAdTypeParameter() string {
	return adType
}
func GetExpiredTimeParameter() string {
	return expireTime
}
func GetAdBucket() string {
	return adBucket
}
func GetImageResourceBucket() string {
	return imageResourceBucket
}

func GetDefaultValue() string {
	return defaultValue
}
func GetOriginalFileDescription() string {
	return originalFileDescription
}
