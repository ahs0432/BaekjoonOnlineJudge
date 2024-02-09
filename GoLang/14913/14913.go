package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, d, k int
	fmt.Fscanln(reader, &a, &d, &k)

	valid := false
	var ans int

	for i := 1; i <= 1000000; i++ {
		if a+(i-1)*d == k {
			valid = true
			ans = i
			break
		}
	}

	if valid {
		fmt.Println(ans)
	} else {
		fmt.Println("X")
	}
}
