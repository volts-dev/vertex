package app

import (
	"encoding/json"
	"sync"

	"github.com/volts-dev/vertex/core/errors"
	"github.com/volts-dev/vertex/core/js"
)

type Storage struct {
	name  string
	mutex sync.RWMutex
}

func init() {
	js.RegisterInterface(GetStorageInterface)
}

var storageinterface js.Value

// GetInterface get the Storage interface
func GetStorageInterface() js.Value {
	singleton.Do(func() {

		if storageinterface = js.Global().Get("Storage"); storageinterface.Error() != nil {
			storageinterface = js.Undefined()
		}

		js.Register(storageinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return storageinterface
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

	DefaultWindow().Get(s.name).Call("setItem", k, string(b))
	return nil
}

func (s *Storage) Get(k string, v any) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	item := DefaultWindow().Get(s.name).Call("getItem", k)
	if item.IsNull() {
		return nil
	}

	itemv, err := item.String()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(itemv), v)
}

func (s *Storage) Del(k string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	DefaultWindow().Get(s.name).Call("removeItem", k)
}

func (s *Storage) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	DefaultWindow().Get(s.name).Call("clear")
}

func (s *Storage) Len() (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.len()
}

func (s *Storage) len() (int, error) {
	return DefaultWindow().Get(s.name).Get("length").Int()
}

func (s *Storage) ForEach(f func(key string)) {
	s.mutex.Lock()
	length, _ := s.len()
	keys := make(map[string]struct{}, length)
	for i := 0; i < length; i++ {
		key := DefaultWindow().Get(s.name).Call("key", i)
		if key.Truthy() {
			keystr, _ := key.String()
			keys[keystr] = struct{}{}
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
	return !DefaultWindow().Get(s.name).Call("getItem", k).IsNull()
}
