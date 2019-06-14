package storage

import (
	"errors"
	"sync"
)

type MemoryStorage struct {
	data sync.Map
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: sync.Map{},
	}
}

func (ms *MemoryStorage) Delete(key string) {
	ms.data.Delete(key)
}

func (ms *MemoryStorage) Fetch(key string) (interface{}, error) {
	data, ok := ms.data.Load(key)

	if !ok {
		return nil, errors.New("key not found")
	}

	return data, nil
}

func (ms *MemoryStorage) Save(key string, value interface{}) {
	ms.data.Store(key, value)
}
