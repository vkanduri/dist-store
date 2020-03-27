package hashstore

import (
	"errors"
	"sync"
)

type inMemStore struct {
	// we can use sync.Map which is a concurrent map
	//mapper sync.Map
	mapper map[string]string
	shardID int
	keyLockers map[string]*sync.Mutex
	creatorLocker sync.Mutex
}

func NewInMemStore(shardID int) inMemStore {
	return inMemStore{
		mapper:  make(map[string]string),
		keyLockers: make(map[string]*sync.Mutex),
		shardID: shardID,
	}
}

func (s inMemStore) Write(key string, value string) error{
	_, ok := s.keyLockers[key]
	if  !ok {
		s.creatorLocker.Lock()
		s.keyLockers[key] = &sync.Mutex{}
		s.creatorLocker.Unlock()
	}
	s.keyLockers[key].Lock()
	s.mapper[key] = value
	s.keyLockers[key].Unlock()
	return nil
}

func (s inMemStore) Read(key string) (string, error){
	v, ok := s.mapper[key]
	if !ok {
		return v, errors.New("key does not exist in the store")
	}
	return v, nil
}

func (s inMemStore) Dump() map[string]string{
	return s.mapper
}

func (s inMemStore)Healthy() bool{
	return true
}