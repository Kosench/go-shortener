package repository

import (
	"fmt"
	"sync"
)

type MemoryRepository struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		urls: make(map[string]string),
	}
}

func (r *MemoryRepository) Save(code, url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.urls[code] = url
	return nil
}

func (r *MemoryRepository) Find(code string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, exists := r.urls[code]
	if !exists {
		return "", ErrNotFound
	}

	return url, nil
}

var ErrNotFound = fmt.Errorf("URL not found")
