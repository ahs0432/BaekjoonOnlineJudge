package main

import (
	"bufio"
	"fmt"
	"os"
)

func changeNum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	now := changeNum(n)
	next := changeNum(changeNum(n))

	for now != 1 && next != 1 {
		if now == next {
			return false
		}

		now = changeNum(now)
		next = changeNum(changeNum(next))
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	if isHappy(n) {
		fmt.Println("HAPPY")
	} else {
		fmt.Println("UNHAPPY")
	}
}
