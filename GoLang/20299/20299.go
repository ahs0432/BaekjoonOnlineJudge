package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, k, l int
	fmt.Fscanln(reader, &n, &k, &l)

	ans := []int{}
	tmp := make([]int, 3)

	var sum int
	var check bool
	for i := 0; i < n; i++ {
		sum = 0
		check = true
		for j := 0; j < 3; j++ {
			fmt.Fscan(reader, &tmp[j])
			if tmp[j] < l {
				check = false
			}
			sum += tmp[j]
		}

		if check && sum >= k {
			ans = append(ans, tmp...)
		}
	}

	fmt.Fprintln(writer, len(ans)/3)
	for i := 0; i < len(ans); i++ {
		fmt.Fprint(writer, ans[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
