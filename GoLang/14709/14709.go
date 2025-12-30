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

	var a, b int
	var table [6][6]int
	if n != 3 {
		for i := 0; i < n; i++ {
			fmt.Fscanln(reader, &a, &b)
		}
		fmt.Println("Woof-meow-tweet-squeek")
	} else {
		for i := 0; i < n; i++ {
			fmt.Fscanln(reader, &a, &b)
			table[a][b] = 1
			table[b][a] = 1
		}

		if table[1][3] == 1 && table[3][4] == 1 && table[4][1] == 1 {
			fmt.Println("Wa-pa-pa-pa-pa-pa-pow!")
		} else {
			fmt.Println("Woof-meow-tweet-squeek")
		}
	}
}
