package utils

var (
	filmCollection           = "Film"
	filmRouteCollection      = "FilmRoute"
	filmPopularityCollection = "FilmPopularity"
	collectionList           = []string{filmCollection, filmRouteCollection, filmPopularityCollection}

	searchFilmIndex = "film"

	elasticsearchUseSSL = false
)

func GetFilmCollection() string {
	return filmCollection
}
func GetFilmPopularityCollection() string {
	return filmPopularityCollection
}
func GetFilmRouteCollection() string {
	return filmRouteCollection
}
func GetCollectionList() []string {
	return collectionList
}

func GetSearchFilmIndex() string {
	return searchFilmIndex
}

func GetElasticsearchUseSSL() bool {
	return elasticsearchUseSSL
}
