package dist_store

import (
	"fmt"
	"gitlab.com/etherlabs/dist-store/pkg/hashstore"
)

type distStore struct {
	storers map[int32]hashstore.Storer
	shardCount int32
}

func NewDistStorer(shardCount int32) distStore{
	ds := distStore{storers: map[int32]hashstore.Storer{}, shardCount: shardCount}
	for i := 0; i < int(shardCount); i++ {
		ds.storers[int32(i)] = hashstore.NewInMemStore(i+1)
	}
	return ds
}


func (d distStore) Write(key string, value string) error {
	shardID := d.hasherIndex(key)
	return d.storers[shardID].Write(key, value)
}

func (d distStore) Read(key string) (string, error) {
	shardID := d.hasherIndex(key)
	return d.storers[shardID].Read(key)
}

func (d distStore) Dump(){
	for i, s := range d.storers{
		fmt.Println("dumping from store: ",i+1)
		fmt.Println(s.Dump())
	}
}

func (d distStore)hasherIndex(key string) int32{
	var keyHash int32
	for i, r := range key {
		keyHash = keyHash + r*(int32(i))
	}
	return keyHash % int32(d.shardCount)
}
