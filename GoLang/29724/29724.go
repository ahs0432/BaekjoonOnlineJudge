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

	box, count := 0, 0
	var t string
	var w, h, l int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &t, &w, &h, &l)
		if t == "A" {
			count += (w / 12) * (h / 12) * (l / 12)
			box += 1000
		} else {
			box += 6000
		}
	}
	fmt.Println(box + (count * 500))
	fmt.Println(count * 4000)
}
