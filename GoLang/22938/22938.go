package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x1, y1, r1, x2, y2, r2 int
	fmt.Fscanln(reader, &x1, &y1, &r1)
	fmt.Fscanln(reader, &x2, &y2, &r2)

	if (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) < (r1+r2)*(r1+r2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
