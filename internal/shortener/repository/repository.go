package repository

import "github.com/Kosench/go-shortener/internal/shortener/model"

type Repository interface {
	Save(url model.URL) error
	Find(code string) (model.URL, error)
}
