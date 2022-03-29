package store

type CacheRepository interface {
	Get(string) (string, error)
	Set(string, string) error

	GetAll() (map[string]string, error)
}
