package structure

import (
	"go-basic-final/internal/file"
	"os"
	"sync"
)

type Synchronization interface {
	Synchronize(mainDir string, subDir string) error
}

type SynchronizationImp struct {
	sync.Mutex
}

func (s *SynchronizationImp) Synchronize(mainDir string, subDir string) error {
	s.Lock()
	defer s.Unlock()

	mainStruct, err := GenStructure(mainDir)
	if err != nil {
		return err
	}

	subStruct, err := GenStructure(subDir)
	if err != nil {
		return err
	}

	if err := syncStructs(mainStruct, subStruct); err != nil {
		return err
	}

	return nil
}

func syncStructs(mainStruct *Item, subStruct *Item) error {
	if err := deleteItems(mainStruct, subStruct); err != nil {
		return err
	}

	if err := copyItems(mainStruct, subStruct); err != nil {
		return err
	}

	return nil
}

func copyItems(mainStruct *Item, subStruct *Item) error {
	for k, item := range mainStruct.Items {
		if _, ok := subStruct.Items[k]; !ok {

			if DIRECTORY == item.ItemType {
				newSubItem := NewItem(DIRECTORY, item.Name)
				subStruct.AddChild(newSubItem)

				if err := os.Mkdir(newSubItem.Path, os.ModePerm); err != nil {
					return err
				}

				err := copyItems(item, newSubItem)
				if err != nil {
					return err
				}
			} else {
				newSubItem := NewItem(FILE, item.Name)
				subStruct.AddChild(newSubItem)
				err := file.CopyFile(item.Path, newSubItem.Path)
				if err != nil {
					return err
				}
			}
		} else {
			if DIRECTORY == item.ItemType {
				err := copyItems(item, subStruct.Items[k])
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func deleteItems(mainStruct *Item, subStruct *Item) error {
	for k, item := range subStruct.Items {
		if _, ok := mainStruct.Items[k]; !ok {
			err := os.RemoveAll(item.Path)
			if err != nil {
				return err
			}
		} else {
			if DIRECTORY == item.ItemType {
				err := deleteItems(mainStruct.Items[k], item)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
