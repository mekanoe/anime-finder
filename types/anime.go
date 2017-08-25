package types

type Anime struct {
	Title  string `json:"title,omitempty"`
	ID     string `json:"id,omitempty"` // ksuid
	Poster string `json:"poster,omitempty"`

	ServiceID struct {
		Kitsu string `json:"kitsu,omitempty"`
	} `json:"service_id,omitempty"`

	Genres []string `json:"genres,omitempty"`
}
