package main

import "fmt"

func main() {
	for {
		var n1, n2 int
		_, err := fmt.Scanln(&n1, &n2)

		if err != nil {
			return
		}

		fmt.Println(n1 + n2)
	}
}
