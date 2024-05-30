package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(num int) int {
	res := 0
	for {
		if num == 0 {
			break
		}
		res += num % 10
		num /= 10
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var l, d, x int
	fmt.Fscanln(reader, &l)
	fmt.Fscanln(reader, &d)
	fmt.Fscanln(reader, &x)

	for i := l; i <= d; i++ {
		if sum(i) == x {
			fmt.Println(i)
			break
		}
	}

	for i := d; i >= l; i-- {
		if sum(i) == x {
			fmt.Println(i)
			break
		}
	}
}
