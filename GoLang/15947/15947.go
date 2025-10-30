package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	count := (n-1)/14 + 2
	pos := (n - 1) % 14

	switch pos {
	case 0, 12:
		fmt.Print("baby")
	case 1, 13:
		fmt.Print("sukhwan")
	case 2, 6, 10:
		if count < 5 {
			fmt.Print("tu")
			for i := 0; i < count; i++ {
				fmt.Print("ru")
			}
		} else {
			fmt.Printf("tu+ru*%d", count)
		}
	case 3, 7, 11:
		if count-1 < 5 {
			fmt.Print("tu")
			for i := 0; i < count-1; i++ {
				fmt.Print("ru")
			}
		} else {
			fmt.Printf("tu+ru*%d", count-1)
		}
	case 4:
		fmt.Print("very")
	case 5:
		fmt.Print("cute")
	case 8:
		fmt.Print("in")
	case 9:
		fmt.Print("bed")
	default:
		break
	}
	fmt.Println()
}
