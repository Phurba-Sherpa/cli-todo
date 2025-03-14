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

func ReadItems(fileName string) ([]Item, error) {
	b, err := os.ReadFile(fileName)

	if err != nil {
		return []Item{}, nil
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, nil
}
