package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	ans := table[n-1]
	if table[2]-table[1] == table[1]-table[0] {
		ans += table[2] - table[1]
	} else if table[2]/table[1] == table[1]/table[0] {
		ans *= (table[2] / table[1])
	}
	fmt.Println(ans)
}
