package store

type Store interface {
	Cache() *CacheRepository
}
