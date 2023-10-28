package model

type Links struct {
	id        uint64
	shortLink string `json:"short_link"`
	longLink  string `json:"long_link"`
}
