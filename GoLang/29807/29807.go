package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(reader, &t)

	table := make([]int, 5)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &table[i])
	}

	ans := 0
	if table[0] > table[2] {
		ans += (table[0] - table[2]) * 508
	} else {
		ans += (table[2] - table[0]) * 108
	}

	if table[1] > table[3] {
		ans += (table[1] - table[3]) * 212
	} else {
		ans += (table[3] - table[1]) * 305
	}
	if table[4] > 0 {
		ans += table[4] * 707
	}
	ans *= 4763

	fmt.Println(ans)
}
