package dist_store

import (
	"fmt"
	"gitlab.com/etherlabs/dist-store/pkg/hashstore"
)

type cyclicDistStore struct {
	storers map[int32]hashstore.Storer
	shardCount int32
}

func NewCyclicDistStorer(shardCount int32) cyclicDistStore {
	ds := cyclicDistStore{storers: map[int32]hashstore.Storer{}, shardCount: shardCount}
	for i := 0; i < int(shardCount); i++ {
		ds.storers[int32(i)] = hashstore.NewInMemStore(i+1)
	}
	return ds
}


func (d cyclicDistStore) Write(key string, value string) error {


	shardID, redShardID := d.hasherIndex(key)
	go d.storers[redShardID].Write(key, value)
	return d.storers[shardID].Write(key, value)
}

func (d cyclicDistStore) Read(key string) (string, error) {
	shardID,redShardID := d.hasherIndex(key)
	if d.storers[shardID].Healthy(){
		return d.storers[shardID].Read(key)
	} else {
		return d.storers[redShardID].Read(key)
	}
}

func (d cyclicDistStore) Dump(){
	for i, s := range d.storers{
		fmt.Println("dumping from store: ",i+1)
		fmt.Println(s.Dump())
	}
}

func (d cyclicDistStore)hasherIndex(key string) (int32, int32){
	var keyHash int32
	for i, r := range key {
		keyHash = keyHash + r*(int32(i))
	}
	sID := keyHash % int32(d.shardCount)
	rsIDOffset := keyHash % 2
	if rsIDOffset == 0 {
		if sID+1 >= d.shardCount{
			return sID, 0
		}
		return sID, sID +1
	}
	if sID-1 < 0{
		return sID, d.shardCount -1
	}
	return sID, sID -1

}
