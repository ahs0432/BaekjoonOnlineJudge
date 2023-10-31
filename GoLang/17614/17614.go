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

	sum := 0
	var tmp int
	for i := 3; i <= n; i++ {
		tmp = i

		for tmp > 0 {
			if tmp%10 == 3 || tmp%10 == 6 || tmp%10 == 9 {
				sum++
			}
			tmp /= 10
		}
	}
	fmt.Println(sum)
}
