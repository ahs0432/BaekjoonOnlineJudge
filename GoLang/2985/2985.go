package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	if a+b == c {
		fmt.Printf("%d+%d=%d\n", a, b, c)
	} else if a-b == c {
		fmt.Printf("%d-%d=%d\n", a, b, c)
	} else if a*b == c {
		fmt.Printf("%d*%d=%d\n", a, b, c)
	} else if a/b == c {
		fmt.Printf("%d/%d=%d\n", a, b, c)
	} else if a == b+c {
		fmt.Printf("%d=%d+%d\n", a, b, c)
	} else if a == b-c {
		fmt.Printf("%d=%d-%d\n", a, b, c)
	} else if a == b*c {
		fmt.Printf("%d=%d*%d\n", a, b, c)
	} else if a == b/c {
		fmt.Printf("%d=%d/%d\n", a, b, c)
	}
}
