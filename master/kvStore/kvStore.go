package kvStore

// needs to implement an interface to a cache 
// specific cache inits can be manually configured

import (
	"sync"
	"errors"
)

// map-based data store with sync mutex for concurrency protection
type KVStore struct {
	sync.Mutex
	map[string]string // maybe changed later for flexibility
}

func New(...) (*KVStore, error) {

}

func (kv *KVStore) Get(key string) (string, error) {

}

func (kv *KVStore) Set(key string, value string) error {

}

func (kv *KVStore) Del(key string) error {

}

// maybe inc?

