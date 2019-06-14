package storage

type Storage interface {
	Delete(key string)
	Fetch(key string) (interface{}, error)
	Save(key string, value interface{})
}
