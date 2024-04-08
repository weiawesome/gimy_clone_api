package service

import "strconv"

const (
	specificFilmType       = "FilmInformation"
	specificFilmRoutesType = "FilmRoutesInformation"
	searchCelebrityType    = "SearchCelebrity"
	searchContentType      = "SearchContent"
	popularFilmTypeType    = "PopularFilmType"
	popularCategoryType    = "PopularCategory"
	filterType             = "FilterFilms"
	rankedType             = "RankedFilms"
	basicType              = "BasicFilms"
)

func FormatKeySpecificFilm(filmId string) string {
	return specificFilmType + "-" + filmId
}
func FormatKeySpecificFilmRoutes(filmId string) string {
	return specificFilmRoutesType + "-" + filmId
}
func FormatKeySearchCelebrity(content string, offset int32, limit int32) string {
	return searchCelebrityType + "-" + content + "-" + strconv.Itoa(int(offset)) + "-" + strconv.Itoa(int(limit))
}
func FormatKeySearchContent(content string, offset int32, limit int32) string {
	return searchContentType + "-" + content + "-" + strconv.Itoa(int(offset)) + "-" + strconv.Itoa(int(limit))
}
func FormatKeyPopularFilmType(filmType string, offset int32, limit int32) string {
	return popularFilmTypeType + "-" + filmType + "-" + strconv.Itoa(int(offset)) + "-" + strconv.Itoa(int(limit))
}
func FormatKeyPopularFilmCategory(category string, offset int32, limit int32) string {
	return popularCategoryType + "-" + category + "-" + strconv.Itoa(int(offset)) + "-" + strconv.Itoa(int(limit))
}
func FormatKeyFilter(filmType string, category string, releaseYear uint32, orderType string, offset int32, limit int32) string {
	return filterType + "-" + filmType + "-" + category + "-" + strconv.Itoa(int(releaseYear)) + "-" + orderType + "-" + strconv.Itoa(int(offset)) + "-" + strconv.Itoa(int(limit))
}
func FormatKeyRanked(category string, offset int32, limit int32) string {
	return rankedType + "-" + category + "-" + strconv.Itoa(int(offset)) + "-" + strconv.Itoa(int(limit))
}
func FormatKeyBasic(offset int32, limit int32) string {
	return basicType + "-" + strconv.Itoa(int(offset)) + "-" + strconv.Itoa(int(limit))
}
