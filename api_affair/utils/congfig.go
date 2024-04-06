package utils

const (
	version = "v1"

	filmId       = "FilmId"
	filmType     = "FilmType"
	filmCategory = "FilmCategory"

	adType = "AdType"

	offset      = "offset"
	limit       = "limit"
	content     = "content"
	category    = "category"
	location    = "location"
	releaseYear = "release_year"
	orderType   = "order_type"

	defaultOffset      = 0
	defaultLimit       = 60
	defaultReleaseYear = 0
	defaultValue       = ""
)

func GetVersion() string {
	return version
}

func GetFilmIdRouteParameter() string {
	return filmId
}
func GetFilmTypeRouteParameter() string {
	return filmType
}
func GetFilmCategoryRouteParameter() string {
	return filmCategory
}
func GetAdTypeRouteParameter() string {
	return adType
}
func GetOffsetParameter() string {
	return offset
}
func GetLimitParameter() string {
	return limit
}
func GetContentParameter() string {
	return content
}
func GetCategoryParameter() string {
	return category
}
func GetLocationParameter() string {
	return location
}
func GetReleaseYearParameter() string {
	return releaseYear
}
func GetOrderTypeParameter() string {
	return orderType
}
func GetDefaultOffset() int32 {
	return defaultOffset
}
func GetDefaultLimit() int32 {
	return defaultLimit
}
func GetDefaultReleaseYear() uint32 {
	return uint32(defaultReleaseYear)
}
func GetDefaultValue() string {
	return defaultValue
}
