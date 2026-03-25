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

	x, y, count, dis := 0, 0, 0, 0
	for {
		dis++
		for i := 0; i < dis; i++ {
			if count == n {
				fmt.Println(x, y)
				return
			}
			count++
			y++
		}

		for i := 0; i < dis; i++ {
			if count == n {
				fmt.Println(x, y)
				return
			}
			count++
			x++
		}

		dis++
		for i := 0; i < dis; i++ {
			if count == n {
				fmt.Println(x, y)
				return
			}
			count++
			y--
		}

		for i := 0; i < dis; i++ {
			if count == n {
				fmt.Println(x, y)
				return
			}
			count++
			x--
		}
	}
}
