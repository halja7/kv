package kv_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/halja7/kv/db"
	"github.com/halja7/kv/kv"
)

type fakePersistenceImpl struct{}

func (p *fakePersistenceImpl) Readall() (map[string]string, error) {
	return map[string]string{}, nil
}

func (p *fakePersistenceImpl) Flush(data map[string]string) error {
	return nil
}

func NewFakePersistence() *fakePersistenceImpl {
	return &fakePersistenceImpl{}
}

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

	tests := []struct {
		name     string
		key      string
		value    string
		expected string
	}{
		{"SetAndGetBasic", "foo", "bar", "bar"},
		{"EmptyStringValue", "foo", "", ""},
		{"OverwriteExisting", "overwrite", "initial", "updated"},
		{"LargeValue", "foo", strings.Repeat("a", 1e6), strings.Repeat("a", 1e6)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store, err := kv.NewStore(NewFakePersistence())
			if err != nil {
				t.Fatalf("Error creating new store: %v", err)
			}

			initial := store.Get(tt.key)
			if initial != "" {
				t.Errorf("Expected empty string, got %q", initial)
			}

			store.Set(tt.key, tt.value)

			if tt.name == "OverwriteExisting" {
				store.Set(tt.key, "updated")
			}
			got := store.Get(tt.key)
			if got != tt.expected {
				t.Errorf("Incorrect value for key %q: got %q, want %q", tt.key, got, tt.expected)
			}
		})
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

	if value != "" {
		t.Errorf("%q unexpectedly non-empty store: ", value)
	}

	store.Set("foo", "bar")

	value = store.Get("foo")

	if value != "bar" {
		t.Errorf("%q is the incorrect value: ", value)
	}

	store2, err := kv.NewStore(db)
	if err != nil {
		t.Error("Error creating new store")
	}

	value = store2.Get("foo")

	fmt.Println("value TestPersist", value)
	if value != "bar" {
		t.Errorf("%q is the incorrect value: ", value)
	}
}
