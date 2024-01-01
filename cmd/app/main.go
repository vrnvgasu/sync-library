package main

import (
	"go-basic-final/internal/structure"
)

func main() {
	err := structure.Synchronization("/home/dmitrii/Загрузки", "/home/dmitrii/test")
	if err != nil {
		panic(err)
	}
}
