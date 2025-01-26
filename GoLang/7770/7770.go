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

	i, ans := 0, 0
	for i <= n {
		ans++
		i += (ans*2-1)*(ans*2-1) - 4*(ans*(ans-1)/2)
	}

	if i > n {
		ans--
	}
	fmt.Println(ans)
}
