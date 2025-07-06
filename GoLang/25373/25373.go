package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscanln(reader, &n)

	if n == 1 {
		fmt.Println(1)
	} else if n <= 3 {
		fmt.Println(2)
	} else if n <= 6 {
		fmt.Println(3)
	} else if n <= 10 {
		fmt.Println(4)
	} else if n <= 15 {
		fmt.Println(5)
	} else if n <= 21 {
		fmt.Println(6)
	} else {
		fmt.Println((n-22)/7 + 7)
	}
}
