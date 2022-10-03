package main

import "fmt"

func main() {
	for {
		var n1, n2 int
		fmt.Scanln(&n1, &n2)
		if n1 == 0 && n2 == 0 {
			break
		}
		fmt.Println(n1 + n2)
	}
}
