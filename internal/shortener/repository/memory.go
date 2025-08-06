package repository

import (
	"fmt"
	"github.com/Kosench/go-shortener/internal/shortener/model"
	"sync"
)

var ErrNotFound = fmt.Errorf("URL not found")
var ErrCodeExists = fmt.Errorf("code already exists")

type MemoryRepository struct {
	mu   sync.RWMutex
	urls map[string]model.URL
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		urls: make(map[string]model.URL),
	}
}

func (r *MemoryRepository) Save(url model.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock() // Добавлен defer для разблокировки
	if _, exists := r.urls[url.Code]; exists {
		return ErrCodeExists
	}
	r.urls[url.Code] = url // Сохраняем весь объект model.URL
	return nil
}

func (r *MemoryRepository) Find(code string) (model.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, exists := r.urls[code]
	if !exists {
		return model.URL{}, ErrNotFound // Возвращаем пустую структуру и ошибку
	}
	return url, nil
}
