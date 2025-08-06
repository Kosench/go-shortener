package repository

type Repository interface {
	Save(code, url string) error
	Find(code string) (string, error)
}
