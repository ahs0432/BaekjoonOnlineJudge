package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		var n1, n2 int
		fmt.Scanln(&n1, &n2)

		fmt.Println("Case #"+strconv.Itoa(i)+":", n1+n2)
	}
}
