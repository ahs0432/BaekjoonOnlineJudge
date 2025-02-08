package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	checkList := []byte{'q', 'w', 'e', 'r', 't', 'a', 's', 'd', 'f', 'g', 'z', 'x', 'c', 'v', 'Q', 'W', 'E', 'R', 'T', 'A', 'S', 'D', 'F', 'G', 'Z', 'X', 'C', 'V'}
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	fmt.Fscanln(reader, &s)

	check := false

	for _, c := range checkList {
		if s[len(s)-1] == c {
			check = true
			break
		}
	}

	if check {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
