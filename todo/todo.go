package todo

import (
	"encoding/json"
	"os"
)

type Item struct {
	Text string
}

func SaveItem(fileName string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fileName, b, 0644); err != nil {
		return err
	}

	return nil
}
