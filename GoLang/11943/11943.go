package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	t := make([]int, 4)
	fmt.Fscanln(reader, &t[0], &t[1])
	fmt.Fscanln(reader, &t[2], &t[3])

	if t[0]+t[3] < t[1]+t[2] {
		fmt.Println(t[0] + t[3])
	} else {
		fmt.Println(t[1] + t[2])
	}
}
