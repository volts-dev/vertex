//go:build js && wasm

package vertex

import (
	"encoding/json"
	"errors"
	"sync"
)

type Storage struct {
	name  string
	mutex sync.RWMutex
}

func newStorage(name string) *Storage {
	return &Storage{name: name}
}

func (s *Storage) Set(k string, v any) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = errors.New("setting storage value failed").
				WithTag("storage-type", s.name).
				WithTag("key", k).
				Wrap(r.(error))
		}
	}()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	Window().Get(s.name).Call("setItem", k, string(b))
	return nil
}

func (s *Storage) Get(k string, v any) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	item := Window().Get(s.name).Call("getItem", k)
	if item.IsNull() {
		return nil
	}

	return json.Unmarshal([]byte(item.String()), v)
}

func (s *Storage) Del(k string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	Window().Get(s.name).Call("removeItem", k)
}

func (s *Storage) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	Window().Get(s.name).Call("clear")
}

func (s *Storage) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.len()
}

func (s *Storage) len() int {
	return Window().Get(s.name).Get("length").Int()
}

func (s *Storage) ForEach(f func(key string)) {
	s.mutex.Lock()
	length := s.len()
	keys := make(map[string]struct{}, length)
	for i := 0; i < length; i++ {
		key := Window().Get(s.name).Call("key", i)
		if key.Truthy() {
			keys[key.String()] = struct{}{}
		}
	}
	s.mutex.Unlock()

	for key := range keys {
		f(key)
	}
}

func (s *Storage) Contains(k string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return !Window().Get(s.name).Call("getItem", k).IsNull()
}
