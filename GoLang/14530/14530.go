package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, y int
	fmt.Fscanln(reader, &x, &y)

	john := x

	dis, way := 1, 1
	now, ans := 1, 0
	for {
		for i := 0; i < dis; i++ {
			john += way
			ans++

			if john == y {
				fmt.Println(ans)
				return
			}
		}
		now *= 2

		dis = abs(john-x) + now
		way *= -1
	}
}
