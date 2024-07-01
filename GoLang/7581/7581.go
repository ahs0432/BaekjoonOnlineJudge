package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var l, w, h, v int

	for {
		fmt.Fscanln(reader, &l, &w, &h, &v)

		if l == 0 && l == w && l == h && l == v {
			break
		}

		if l == 0 {
			l = v / w / h
		} else if w == 0 {
			w = v / l / h
		} else if h == 0 {
			h = v / l / w
		} else if v == 0 {
			v = l * w * h
		}
		fmt.Println(l, w, h, v)
	}
}
