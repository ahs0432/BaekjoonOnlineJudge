package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var e, f, c int
	fmt.Fscanln(reader, &e, &f, &c)

	n := (e+f)/c + (e+f)%c
	ans := (e + f) / c

	for n/c != 0 {
		ans += n / c
		n = n/c + n%c
	}
	fmt.Println(ans)
}
