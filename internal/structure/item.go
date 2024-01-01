package structure

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
)

const (
	DIRECTORY = iota
	FILE
)

type Item struct {
	ItemType int
	Name     string
	Items    map[string]*Item
	Parent   *Item
	Path     string
}

func (i *Item) Hash() string {
	h := sha1.New()
	h.Write([]byte(i.Name + strconv.Itoa(i.ItemType)))
	return hex.EncodeToString(h.Sum(nil))
}

func NewItem(itemType int, name string) *Item {
	switch itemType {
	case DIRECTORY:
	case FILE:
	default:
		panic(fmt.Sprintf("itemType %d is not exist", itemType))
	}

	return &Item{itemType, name, make(map[string]*Item), nil, name}
}

func (i *Item) AddChild(item *Item) {
	if _, ok := i.Items[item.Hash()]; ok {
		return
	}

	item.Path = i.Path + "/" + item.Path
	item.Parent = i

	i.Items[item.Hash()] = item
}
