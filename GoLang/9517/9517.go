package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k int
	fmt.Fscanln(reader, &k)

	var n, time, t int
	var z string
	fmt.Fscanln(reader, &n)

	time = 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &t, &z)

		if z == "T" {
			time += t
			if time < 210 {
				k++
				if k == 9 {
					k = 1
				}
			} else {
				fmt.Println(k)
				break
			}
		} else {
			time += t
			if time >= 210 {
				fmt.Println(k)
				break
			}
		}
	}
}
