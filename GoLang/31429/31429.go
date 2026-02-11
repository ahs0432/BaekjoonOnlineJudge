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

	switch n {
	case 1:
		fmt.Println("12 1600")
	case 2:
		fmt.Println("11 894")
	case 3:
		fmt.Println("11 1327")
	case 4:
		fmt.Println("10 1311")
	case 5:
		fmt.Println("9 1004")
	case 6:
		fmt.Println("9 1178")
	case 7:
		fmt.Println("9 1357")
	case 8:
		fmt.Println("8 837")
	case 9:
		fmt.Println("7 1055")
	case 10:
		fmt.Println("6 556")
	case 11:
		fmt.Println("6 773")
	}
}
