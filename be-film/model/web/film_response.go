package web

type FilmResponse struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Genre    []string `json:"genre"`
	Rating   string   `json:"rating"`
	Duration string   `json:"duration"`
	Quality  string   `json:"quality"`
	Trailer  string   `json:"trailer"`
	Watch    string   `json:"watch"`
}
