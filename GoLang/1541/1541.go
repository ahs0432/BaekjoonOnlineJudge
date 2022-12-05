package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscanln(reader, &n)

	total := 0

	numTable := []int{}

	tmp := 0
	now := 0
	minus := false
	for _, num := range n {
		if num == '+' {
			now += tmp
			tmp = 0
		} else if num == '-' {
			now += tmp
			if minus {
				numTable = append(numTable, -now)
			} else {
				numTable = append(numTable, now)
			}
			minus = true
			tmp = 0
			now = 0
		} else {
			tmp *= 10
			tmp += (int(num) - 48)
		}
	}

	now += tmp
	if minus {
		numTable = append(numTable, -now)
	} else {
		numTable = append(numTable, now)
	}

	for _, num := range numTable {
		total += num
	}

	fmt.Println(total)
}
