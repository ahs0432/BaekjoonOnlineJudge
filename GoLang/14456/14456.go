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

	var a, b int
	count1, count2 := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)
		if (a == 1 && b == 2) || (a == 2 && b == 3) || (a == 3 && b == 1) {
			count1++
		} else if (a == 1 && b == 3) || (a == 3 && b == 2) || (a == 2 && b == 1) {
			count2++
		}
	}
	if count1 > count2 {
		fmt.Println(count1)
	} else {
		fmt.Println(count2)
	}
}
