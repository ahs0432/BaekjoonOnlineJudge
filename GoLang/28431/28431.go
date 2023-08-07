package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := make([]int, 0)

	for i := 0; i < 5; i++ {
		var n int
		check := false
		fmt.Fscanln(reader, &n)

		for j := 0; j < len(a); j++ {
			if a[j] == n {
				a = append(a[:j], a[j+1:]...)
				check = true
				break
			}
		}

		if !check {
			a = append(a, n)
		}
	}

	fmt.Println(a[0])
}
