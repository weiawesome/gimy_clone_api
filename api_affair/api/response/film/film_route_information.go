package film

type FilmRouteInformation struct {
	FilmRoutes []FilmRoute `json:"film_routes"`
}
type FilmRoute struct {
	Route    string   `json:"route"`
	Episodes []string `json:"episodes"`
}
