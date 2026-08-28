package kvStore

// needs to implement an interface to a cache 
// specific cache inits can be manually configured

import (
	"sync"
	// "errors"
)

// map-based data store with sync mutex for concurrency protection
type KVStore struct {
	sync.Mutex
	store map[string]string // maybe changed later for flexibility
}

func New() (*KVStore) {
	kv := KVStore{store: map[string]string{}}
	return &kv
}

func (kv *KVStore) Get(key string) (string, bool) { // value string, found bool
	kv.Lock()
	defer kv.Unlock()

	val := kv.store[key]
	if val == "" {
		return val, false
	}
	return val, true
}

func (kv *KVStore) Set(key string, value string) {
	kv.Lock()
	defer kv.Unlock()

	kv.store[key] = value
}

func (kv *KVStore) Del(key string) {
	kv.Lock()
	defer kv.Unlock()

	delete(kv.store, key)
}

// maybe inc?

