package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	ans := 0
	for i := 0; i < n; i++ {
		ans += 1

		if strings.Contains(strconv.Itoa(i), "50") {
			ans++
		}
	}
	fmt.Println(ans)
}
