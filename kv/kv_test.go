package kv_test

import (
  "testing"
  "fmt"
  "os"
  "encoding/json"

  "github.com/halja7/kv/kv"
  "github.com/halja7/kv/db"
)

func TestMain(m *testing.M) {
    code := m.Run() 
    data, err := json.Marshal(&map[string]string{})
    if err != nil {
      fmt.Println("Error clearing out test store")
      os.Exit(1)
    }
    os.WriteFile("../store.json", data, 0644)
    os.Exit(code)
}

func TestKv(t *testing.T) {
  t.Parallel()

  // for now default store is usable
  store, err:= kv.NewStore(nil)
  if err != nil {
    t.Error("error creating new store")
  }

  value := store.Get("foo")
  // assert that value is nil
  if value != ""  {
    t.Errorf("%q unexpectedly non-empty store: ", value)
  }

  store.Set("foo", "bar")

  value = store.Get("foo")

  if value != "bar" {
    t.Errorf("%q is the incorrect value: ", value)
  }
}

func TestPersist(t *testing.T) {
  t.Parallel()

  db := db.NewDb("../store.json")

  store, err := kv.NewStore(db)
  if err != nil {
    t.Error("Error creating new store")
  }

  value := store.Get("foo")

  if value != ""  {
    t.Errorf("%q unexpectedly non-empty store: ", value)
  }

  store.Set("foo", "bar")

  value = store.Get("foo")

  if value != "bar" {
    t.Errorf("%q is the incorrect value: ", value)
  }

  store2, err:= kv.NewStore(db)
  if err != nil {
    t.Error("Error creating new store")
  }

  value = store2.Get("foo")

  fmt.Println("value TestPersist", value);
  if value != "bar" {
    t.Errorf("%q is the incorrect value: ", value)
  }
}











