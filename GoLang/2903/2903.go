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

	cnt := 2
	for ; n > 0; n-- {
		cnt += (cnt - 1)
	}

	fmt.Println(cnt * cnt)
}
