package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := make([]int, 50)
	for i := 0; i < 50; i++ {
		if i == 49 {
			fmt.Fscanln(reader, &table[i])
			break
		}
		fmt.Fscan(reader, &table[i])
	}

	var my int
	fmt.Fscanln(reader, &my)

	for i := 0; i < 50; i++ {
		if table[i] == my {
			if i < 5 {
				fmt.Println("A+")
			} else if i < 15 {
				fmt.Println("A0")
			} else if i < 30 {
				fmt.Println("B+")
			} else if i < 35 {
				fmt.Println("B0")
			} else if i < 45 {
				fmt.Println("C+")
			} else if i < 48 {
				fmt.Println("C0")
			} else {
				fmt.Println("F")
			}

			break
		}
	}
}
