package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var h, m, s, q int
	fmt.Fscanln(reader, &h, &m, &s)
	fmt.Fscanln(reader, &q)
	sec := h*3600 + m*60 + s

	var t, c int
	for ; q > 0; q-- {
		fmt.Fscan(reader, &t)

		if t == 1 {
			fmt.Fscan(reader, &c)
			sec = (sec + c) % 86400
		} else if t == 2 {
			fmt.Fscan(reader, &c)
			sec = (sec - c) % 86400
			if sec < 0 {
				sec += 86400
			}
		} else {
			h = sec / 3600
			m = sec / 60 % 60
			fmt.Fprintln(writer, h, m, sec%60)
		}
	}
	writer.Flush()
}
