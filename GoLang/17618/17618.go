package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	ans := 0
	for i := 1; i <= n; i++ {
		str := strconv.Itoa(i)
		digit := 0

		for _, c := range str {
			digit += int(c) - 48
		}

		if i%digit == 0 {
			ans++
		}
	}

	fmt.Println(ans)
}
