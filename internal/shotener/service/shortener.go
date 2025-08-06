package service

import (
	"encoding/base64"
	"github.com/Kosench/go-shortener/internal/shotener/repository"
	"math/rand"
)

type ShortenerService struct {
	repo repository.Repository
}

func NewShortenerService(repo repository.Repository) *ShortenerService {
	return &ShortenerService{repo: repo}
}

func (s *ShortenerService) Shorten(url string) (string, error) {
	code := generateCode()
	err := s.repo.Save(code, url)
	return code, err
}

func generateCode() string {
	buf := make([]byte, 6)
	rand.Read(buf)
	return base64.URLEncoding.EncodeToString(buf)[:6]
}
