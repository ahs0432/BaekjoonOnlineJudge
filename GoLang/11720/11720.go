package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var num string
	fmt.Scanln(&num)

	total := 0
	for i := 0; i < n; i++ {
		total += (int(num[i]) - 48)
	}
	fmt.Println(total)
}
