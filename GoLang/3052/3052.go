package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var count []int

	for i := 0; i < 10; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		num = num % 42

		if count == nil {
			count = []int{num}
		} else {
			check := true
			for _, n := range count {
				if n == num {
					check = false
					break
				}
			}

			if check {
				count = append(count, num)
			}
		}
	}

	fmt.Fprint(writer, len(count))
}
