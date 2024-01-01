package structure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewItem(t *testing.T) {
	req := require.New(t)

	t.Run("create file", func(t *testing.T) {
		item := NewItem(FILE, "file")
		req.Equal(item.Name, "file")
		req.Equal(item.ItemType, FILE)
		req.Equal(item.Path, "file")
		req.Equal(len(item.Items), 0)
		req.Nil(item.Parent)
	})
	t.Run("create dir", func(t *testing.T) {
		item := NewItem(DIRECTORY, "dir")
		req.Equal(item.Name, "dir")
		req.Equal(item.ItemType, DIRECTORY)
		req.Equal(item.Path, "dir")
		req.Equal(len(item.Items), 0)
		req.Nil(item.Parent)
	})
}

func TestHash(t *testing.T) {
	req := require.New(t)
	item1 := NewItem(FILE, "file")
	item2 := NewItem(FILE, "file2")

	req.Equal(item1.Hash(), item1.Hash())
	req.NotEqual(item1.Hash(), item2.Hash())
}

func TestAddChild(t *testing.T) {
	req := require.New(t)
	dir := NewItem(DIRECTORY, "dir")
	file := NewItem(FILE, "file")
	req.Equal(len(dir.Items), 0)

	dir.AddChild(file)
	req.Equal(len(dir.Items), 1)
	req.Equal(dir, file.Parent)

}
