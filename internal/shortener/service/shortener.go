package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Kosench/go-shortener/internal/shortener/model"
	"github.com/Kosench/go-shortener/internal/shortener/repository"
	"time"
)

const maxRetries = 5

type ShortenerService struct {
	repo repository.Repository
}

func NewShortenerService(repo repository.Repository) *ShortenerService {
	return &ShortenerService{repo: repo}
}

func (s *ShortenerService) Shorten(originalURL string) (string, error) {
	for i := 0; i < maxRetries; i++ {
		code, err := generateCode()
		if err != nil {
			return "", fmt.Errorf("failed to generate code: %w", err)
		}

		urlModel := model.URL{
			Code:      code,
			Original:  originalURL,
			CreatedAt: time.Now().UTC(), // Заполняем время создания
		}

		err = s.repo.Save(urlModel)
		if err != nil {
			if errors.Is(err, repository.ErrCodeExists) {
				continue
			}
			return "", fmt.Errorf("failed to save URL: %w", err)
		}
		// Успешно сохранено
		return code, nil
	}
	return "", fmt.Errorf("failed to generate unique code after %d attempts", maxRetries)
}

func (s *ShortenerService) GetOriginalURL(code string) (model.URL, error) {
	return s.repo.Find(code)
}

func generateCode() (string, error) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buf)[:6], nil
}
