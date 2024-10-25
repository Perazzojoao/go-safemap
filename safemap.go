package safemap

import "sync"

type Map[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V)
	Delete(key K)
	Len() int
	ForEach(f func(K, V))
}

type safeMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

func New[K comparable, V any]() *safeMap[K, V] {
	return &safeMap[K, V]{
		m: make(map[K]V),
	}
}

func (sm *safeMap[K, V]) Get(key K) (V, bool) {
	sm.RLock()
	value, ok := sm.m[key]
	sm.RUnlock()
	return value, ok
}

func (sm *safeMap[K, V]) Set(key K, value V) {
	sm.Lock()
	sm.m[key] = value
	sm.Unlock()
}

func (sm *safeMap[K, V]) Delete(key K) {
	sm.Lock()
	delete(sm.m, key)
	sm.Unlock()
}

func (sm *safeMap[K, V]) Len() int {
	sm.RLock()
	length := len(sm.m)
	sm.RUnlock()
	return length
}

func (sm *safeMap[K, V]) ForEach(f func(K, V)) {
	sm.RLock()
	for k, v := range sm.m {
		f(k, v)
	}
	sm.RUnlock()
}
