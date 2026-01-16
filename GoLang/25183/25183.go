package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	fmt.Fscanln(reader, &s)

	count := 1
	var check bool
	for i := 0; i < len(s)-1; i++ {
		if abs(int(s[i])-int(s[i+1])) == 1 {
			count++
		} else {
			count = 1
		}

		if count == 5 {
			check = true
			break
		}
	}

	if check {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
