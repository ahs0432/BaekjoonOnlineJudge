package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t1, t2 int
	fmt.Fscanln(reader, &t1)
	fmt.Fscanln(reader, &t2)

	count := 2

	for t1-t2 >= 0 {
		t1, t2 = t2, t1-t2
		count++
	}
	fmt.Println(count)
}
