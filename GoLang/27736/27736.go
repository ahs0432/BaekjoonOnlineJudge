package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	total := 0
	var tmp int
	var inv int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		if tmp == 0 {
			inv++
		} else {
			total += tmp
		}
	}

	if (n/2)+(n%2) <= inv {
		fmt.Println("INVALID")
	} else {
		if total > 0 {
			fmt.Println("APPROVED")
		} else {
			fmt.Println("REJECTED")
		}
	}
}
