package film

type Films struct {
	Films []Film `json:"films"`
}
type Film struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Resource string   `json:"resource"`
	State    string   `json:"state"`
	Actors   []string `json:"actors"`
}
