package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x, n, p int
	fmt.Fscanln(reader, &x)
	fmt.Fscanln(reader, &n)

	ans := x
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &p)
		ans += (x - p)
	}

	fmt.Println(ans)
}
