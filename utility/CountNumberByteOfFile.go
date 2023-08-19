package utility

import (
	"bufio"
	"fmt"
	"os"
)

func CountNBOfile(path string) int {
	count := 0
	fh, err := os.Open(path)
	if err != nil {
		fmt.Println("can't open the file", err)
	}
	reader := bufio.NewReader(fh)
	for {
		line, _ := reader.ReadByte()
		if line == 0 {
			break
		}
		count++

	}
	return count
}
