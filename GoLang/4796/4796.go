package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var l, p, v, tmp int
	for i := 1; ; i++ {
		fmt.Fscanln(reader, &l, &p, &v)
		if l == 0 && p == 0 && v == 0 {
			break
		}
		tmp = v % p
		if l < tmp {
			tmp = l
		}
		fmt.Fprintf(writer, "Case %d: %d\n", i, (l*(v/p) + tmp))
	}
	writer.Flush()
}
