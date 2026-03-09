package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var h, x int
	fmt.Fscanln(reader, &h, &x)

	var tmp int
	ans, count := 0, 1

	for i := 0; i < h; i++ {
		fmt.Fscanln(reader, &tmp)
		count *= x
		count %= 1000000007
		ans += tmp * count
		ans %= 1000000007
	}
	fmt.Println(ans)
}
