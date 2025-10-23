package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, l, k int
	fmt.Fscanln(reader, &n, &l, &k)

	var sub1, sub2 int
	e, h := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &sub1, &sub2)
		if sub2 <= l {
			h++
		} else if sub1 <= l {
			e++
		}
	}

	result := min(k, h) * 140
	if k > h {
		result += min(k-h, e) * 100
	}
	fmt.Println(result)
}
