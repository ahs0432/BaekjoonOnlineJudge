package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var c, w, l, p int
	var ans, rep int
	for {
		fmt.Fscanln(reader, &c, &w, &l, &p)
		if c == 0 {
			break
		}

		if c == 1 {
			fmt.Fprintln(writer, 1)
		} else {
			ans = 1
			rep = w * l * p

			for ; rep != 0; rep-- {
				ans *= c
			}
			fmt.Fprintln(writer, ans)
		}
	}
	writer.Flush()
}
