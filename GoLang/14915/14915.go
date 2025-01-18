package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	num := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	var m, n int
	fmt.Fscanln(reader, &m, &n)

	ans := ""
	if m == 0 {
		ans = "0"
	} else {
		for m != 0 {
			ans = num[m%n] + ans
			m /= n
		}
	}

	fmt.Println(ans)
}
