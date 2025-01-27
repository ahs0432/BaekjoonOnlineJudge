package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int64
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	sum := a * b * c

	ans := make([]int, 10)
	for sum != 0 {
		ans[int(sum%10)]++
		sum /= 10
	}

	for i := 0; i < 10; i++ {
		fmt.Println(ans[i])
	}
}
