package file

import (
	"os"
	"testing"
)

func BenchmarkCopy(b *testing.B) {
	testInput, _ := os.Create("./testInput.txt")
	defer os.Remove("./testInput.txt")
	defer testInput.Close()

	for i := 0; i < b.N; i++ {
		CopyFile("./testInput.txt", "./testOutput.txt")
	}

	defer os.Remove("./testOutput.txt")
}
