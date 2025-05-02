package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int64
	fmt.Fscanln(reader, &t)

	var n, tmp int64
	var check bool
	var checkNums []bool
	for i := int64(1); i <= t; i++ {
		fmt.Fscanln(reader, &n)

		if n == 0 {
			fmt.Fprintf(writer, "Case #%d: INSOMNIA\n", i)
			continue
		}

		checkNums = make([]bool, 10)
		for j := int64(1); ; j++ {
			tmp = n * j

			for tmp != 0 {
				checkNums[tmp%10] = true
				tmp /= 10

				check = true
				for _, checkNum := range checkNums {
					if !checkNum {
						check = false
						continue
					}
				}

				if check {
					fmt.Fprintf(writer, "Case #%d: %d\n", i, n*j)
					break
				}
			}

			if check {
				break
			}
		}

	}
	writer.Flush()
}
