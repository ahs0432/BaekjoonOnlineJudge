package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	xi := make([]int, 5)
	yi := make([]int, 5)
	fmt.Fscanln(reader, &xi[0], &xi[1], &xi[2], &xi[3], &xi[4])
	fmt.Fscanln(reader, &yi[0], &yi[1], &yi[2], &yi[3], &yi[4])

	for i := 0; i < 5; i++ {
		if xi[i] == yi[i] {
			fmt.Println("N")
			return
		}
	}
	fmt.Println("Y")
}
