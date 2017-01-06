package storage

import (
    "sync"
    "errors"
)

//TODO REST server
type mStore struct {
    mux sync.RWMutex
    data map[string][]byte
}

var s = newStore()

func newStore() *mStore {
    return &mStore{data: make(map[string][]byte)}
}

func Retrieve(id string) ([]byte, error) {
   s.mux.RLock()
   defer s.mux.RUnlock()

   if e, ok := s.data[id]; ok {
        return e, nil
   }

   return nil, errors.New("Id not found in store!")
}

func Store(id string, encrypted []byte) {
   s.mux.Lock()
   defer s.mux.Unlock()

   s.data[id] = encrypted
}


