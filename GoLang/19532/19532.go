package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d, e, f int
	fmt.Fscanln(reader, &a, &b, &c, &d, &e, &f)
	fmt.Println((c*e-b*f)/(a*e-b*d), (a*f-d*c)/(a*e-b*d))
}
