package todo

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
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
	item1 := s[i]
	item2 := s[j]

	// ** first prio done
	if item1.Done != item2.Done {
		return item1.Done
	}

	// ** second prio 'prio'
	if item1.Priority != item2.Priority {
		return item1.Priority < item2.Priority
	}

	// ** last prio: "index"
	return item1.position < item2.position

}

// ** Opt
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

	// Check if dir exist if not create one
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0644); err != nil {
		log.Fatalln(dir, "Error creating dir\n", err)
	}

	// Try reading file
	b, err := os.ReadFile(fileName)

	if err != nil {
		// If file not found initialize with []
		if os.IsNotExist(err) {
			if err := SaveItem(fileName, []Item{}); err != nil {
				log.Fatalln(fileName, "Error creating file", err)
			}
			return []Item{}, nil
		} else {
			log.Fatalln(fileName, "Error reading content \n", err)
		}
		return nil, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}
