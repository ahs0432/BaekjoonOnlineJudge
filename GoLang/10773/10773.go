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

	list := make([]int, 0)
	total := 0

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)

		if tmp == 0 {
			total -= list[len(list)-1]
			list = list[:len(list)-1]
		} else {
			total += tmp
			list = append(list, tmp)
		}
	}

	fmt.Println(total)
}
