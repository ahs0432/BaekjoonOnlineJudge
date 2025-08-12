package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, t, s, z, d int64
	fmt.Fscanln(reader, &n, &m)
	fmt.Fscanln(reader, &t, &s)

	z = n + (n-1)/8*s
	d = t + m + (m-1)/8*(2*t+s)

	if z < d {
		fmt.Println("Zip")
		fmt.Println(z)
	} else {
		fmt.Println("Dok")
		fmt.Println(d)
	}
}
