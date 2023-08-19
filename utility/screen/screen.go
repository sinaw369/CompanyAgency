package screen

import "fmt"

func Moveto(x, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}
func ClearScreen() {
	fmt.Print("\033[H")
	fmt.Print("\033[2J") //Clear screen
}
