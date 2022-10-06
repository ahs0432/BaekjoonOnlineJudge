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
		total := i
		tmp := total

		for 0 < tmp {
			total += (tmp % 10)
			tmp /= 10
		}

		if total == num {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(0)
}
