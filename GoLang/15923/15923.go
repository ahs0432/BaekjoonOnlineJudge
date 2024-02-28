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

	minx, miny := 2147483647, 2147483647
	maxx, maxy := 0, 0

	var tmpx, tmpy int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmpx, &tmpy)
		if minx > tmpx {
			minx = tmpx
		}

		if maxx < tmpx {
			maxx = tmpx
		}

		if miny > tmpy {
			miny = tmpy
		}

		if maxy < tmpy {
			maxy = tmpy
		}
	}

	fmt.Println((maxx - minx + maxy - miny) * 2)
}
