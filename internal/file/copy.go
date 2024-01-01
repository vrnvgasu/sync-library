package file

import (
	"bufio"
	"fmt"
	"os"
)

func CopyFile(file1 string, file2 string) error {
	source, err := os.Open(file1)
	if err != nil {
		return err
	}
	defer func() { source.Close() }()

	output, err := os.Create(file2)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() { output.Close() }()

	if err := readAndWriteByLine(source, output); err != nil {
		return err
	}

	return nil
}

func readAndWriteByLine(source *os.File, output *os.File) error {
	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		if _, err := fmt.Fprintln(output, scanner.Text()); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
