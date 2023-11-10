package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var c, k, p int
	fmt.Fscanln(reader, &c, &k, &p)

	ans := 0
	for i := 1; i <= c; i++ {
		ans += k*i + p*(i*i)
	}
	fmt.Println(ans)
}
