package main

import "fmt"

func main() {
	var total, n int
	fmt.Scanln(&total)
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var n1, n2 int
		fmt.Scanln(&n1, &n2)
		total -= (n1 * n2)
	}

	if total == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
