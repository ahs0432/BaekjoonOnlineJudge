package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var sx, ex, sy, ey int
	fmt.Fscanln(reader, &sx, &sy)
	fmt.Fscanln(reader, &ex, &ey)

	if n == 1 || m == 1 {
		if sx == ex && sy == ey {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else if sx%2 == sy%2 {
		if ex%2 == ey%2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		if ex%2 != ey%2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
