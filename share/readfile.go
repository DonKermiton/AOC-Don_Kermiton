package share

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("[%s] while opening %s", err, path)
		panic("error")
	}

	reader := bufio.NewScanner(file)

	return reader, nil
}
