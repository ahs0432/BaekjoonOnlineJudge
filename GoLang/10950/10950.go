package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var n1, n2 int
		fmt.Scanln(&n1, &n2)
		fmt.Println(n1 + n2)
	}
}
