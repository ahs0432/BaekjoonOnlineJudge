package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscanln(reader, &num)

	for i := 0; i < num; i++ {
		var n1, n2, n3 int
		fmt.Fscanln(reader, &n1, &n2, &n3)

		answer := 0

		if n3%n1 == 0 {
			answer += (n1 * 100) + (n3 / n1)
		} else {
			answer += (n3 % n1 * 100) + (n3/n1 + 1)
		}

		fmt.Println(answer)
	}
}
