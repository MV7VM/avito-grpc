package model

type Links struct {
	id        uint64
	ShortLink string `json:"shortLink"`
	LongLink  string `json:"longLink"`
}
