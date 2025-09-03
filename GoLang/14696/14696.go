package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp, tmpNum int
	var tmpA, tmpB []int
	for i := 0; i < n; i++ {
		tmpA = make([]int, 5)
		tmpB = make([]int, 5)

		fmt.Fscan(reader, &tmp)
		for j := 0; j < tmp; j++ {
			if j == tmp-1 {
				fmt.Fscanln(reader, &tmpNum)
			} else {
				fmt.Fscan(reader, &tmpNum)
			}
			tmpA[tmpNum]++
		}

		fmt.Fscan(reader, &tmp)
		for j := 0; j < tmp; j++ {
			if j == tmp-1 {
				fmt.Fscanln(reader, &tmpNum)
			} else {
				fmt.Fscan(reader, &tmpNum)
			}
			tmpB[tmpNum]++
		}

		for j := 4; j >= 1; j-- {
			if tmpA[j] > tmpB[j] {
				fmt.Fprintln(writer, "A")
				break
			} else if tmpA[j] < tmpB[j] {
				fmt.Fprintln(writer, "B")
				break
			} else if j == 1 {
				fmt.Fprintln(writer, "D")
			}
		}
	}
	writer.Flush()
}
