package todo

import (
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

// ByPri implements sort.Sort (convention By<name>)
// Need to implement 3 func Len, Swap, Less
type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}
	return s[i].Priority < s[j].Priority
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
		break
	default:
		i.Priority = 2

	}
}

func (i *Item) PrettyD() string {
	if i.Done {
		return "✓"
	}
	return ""
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}

	if i.Priority == 3 {
		return "(3)"
	}

	return ""
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
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

	for i, _ := range items {
		items[i].position = i + 1
	}

	return items, nil
}
