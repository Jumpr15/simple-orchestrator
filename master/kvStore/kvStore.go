package kvStore

// needs to implement an interface to a cache 
// specific cache inits can be manually configured

import (
	"sync"
)

// map-based data store with sync mutex for concurrency protection
type KvStore struct {
	sync.Mutex
	store map[string]any // maybe changed later for flexibility
}

func New() (*KvStore) {
	kv := KvStore{store: map[string]any{}}
	return &kv
}

func (kv *KvStore) Get(key string) (any, bool) { // value string, found bool
	kv.Lock()
	defer kv.Unlock()

	val, ok := kv.store[key]
	if ok {
		return val, true
	}
	return val, false
}

func (kv *KvStore) GetAll() map[string]any {
	kv.Lock()
	defer kv.Unlock()

	return kv.store
}

func (kv *KvStore) Set(key string, value any) {
	kv.Lock()
	defer kv.Unlock()

	kv.store[key] = value
}

func (kv *KvStore) Del(key string) {
	kv.Lock()
	defer kv.Unlock()

	delete(kv.store, key)
}

// maybe inc?

