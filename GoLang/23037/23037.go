package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	now := 0
	sum := 0
	for _, c := range s {
		now = int(c) - 48
		sum += (now * now * now * now * now)
	}
	fmt.Println(sum)
}
