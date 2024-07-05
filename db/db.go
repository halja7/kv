package db

import (
	"encoding/json"
	"fmt"
	"os"
)

// is it typical to put methods on an empty struct?
type Db struct {
	path string
}

func NewDb(filepath string) *Db {
	return &Db{
		path: filepath,
	}
}

func (db *Db) Flush(data map[string]string) error {
	encodedData, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	err = os.WriteFile(db.path, encodedData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) Readall() (map[string]string, error) {
	f, err := os.ReadFile(db.path)
	if err != nil {
		data, _ := json.Marshal(map[string]string{})
		err = os.WriteFile(db.path, data, 0644)
		if err != nil {
			return nil, err
		}

		return map[string]string{}, nil
	}

	var data map[string]string

	err = json.Unmarshal(f, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
