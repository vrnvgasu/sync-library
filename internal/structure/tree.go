package structure

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GenStructure(root string) (*Item, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	if os.IsNotExist(err) {
		return nil, fmt.Errorf("does not exist, %s", root)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("is not a directory, %s", root)
	}

	rootItem := NewItem(DIRECTORY, root)
	setChildren(root, rootItem)

	var files []string
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return rootItem, nil
}

func setChildren(path string, parent *Item) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		var itemType int
		if e.IsDir() {
			itemType = DIRECTORY
		} else {
			itemType = FILE
		}

		child := NewItem(itemType, e.Name())
		parent.AddChild(child)

		if e.IsDir() {
			setChildren(path+"/"+e.Name(), child)
		}
	}
}
