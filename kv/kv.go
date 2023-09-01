package kv

import (
  "fmt"
)

type store struct {
  Data map[string]string
  Db Persistence
}

type Persistence interface {
  Flush(map[string]string) error
  Readall() (map[string]string, error)
}

func NewStore(p Persistence) (*store, error) {
  var data map[string]string

  if p != nil {
    var err error
    data, err = p.Readall()
    if err != nil {
      return nil, err
    }
  }

  if data == nil {
    data = map[string]string{}
  }

  return &store{
    Data: data,
    Db: p,
  }, nil
}

func (s *store) Set(key, value string) { 
  s.Data[key] = value
  if s.Db != nil {
    err := s.Db.Flush(s.Data)
    if err != nil {
      fmt.Printf("Error flushing to disk %v\n", err)
    }
  }
}

func (s *store) Get(key string) string { 
  return s.Data[key]
}

