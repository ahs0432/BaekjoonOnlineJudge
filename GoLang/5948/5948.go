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

	ans := 0
	table := make([]int, 10000)
	for table[n] == 0 {
		table[n] = 1
		n = n/100%10*10 + n/10%10
		n *= n
		ans++
	}
	fmt.Println(ans)
}
