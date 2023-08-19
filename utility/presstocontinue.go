package utility

import (
	"bufio"
	"fmt"
	"os"
)

func PressToContinue() {
	fmt.Print("Press 'Enter' to continue...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		return
	}

}
