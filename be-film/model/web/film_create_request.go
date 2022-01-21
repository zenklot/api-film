package web

type FilmCreateRequest struct {
	Title    string   `json:"title" validate:"required"`
	Genre    []string `json:"genre" validate:"required"`
	Rating   string   `json:"rating" validate:"required"`
	Duration string   `json:"duration" validate:"required"`
	Quality  string   `json:"quality" validate:"required"`
	Trailer  string   `json:"trailer" validate:"required"`
	Watch    string   `json:"watch" validate:"required"`
}
