package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, p int64
	fmt.Fscanln(reader, &n, &p)

	ans := int64(1)
	for i := int64(1); i <= n; i++ {
		ans = (ans * (i % p)) % p
	}
	fmt.Println(ans)
}
