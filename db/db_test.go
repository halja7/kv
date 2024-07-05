package db_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

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

func TestDb(t *testing.T) {
	data := map[string]string{
		"foo":    "bar",
		"baz":    "qux",
		"quux":   "corge",
		"grault": "garply",
		"waldo":  "22",
	}

	db := db.NewDb("../store.json")
	db.Flush(data)
	readData, err := db.Readall()

	if err != nil {
		t.Errorf("db.Readall() returned nil. Expected map[string]string")
	}

	if readData["foo"] != "bar" {
		t.Errorf("Expected %q, got %q", "bar", readData["foo"])
	}

	if readData["waldo"] != "22" {
		t.Errorf("Expected %q, got %q", "22", readData["waldo"])
	}
}
