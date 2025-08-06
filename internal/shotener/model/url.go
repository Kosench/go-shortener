package model

type URL struct {
	Code     string `json:"code"`
	Original string `json:"original"`
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
