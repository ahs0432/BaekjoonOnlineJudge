package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanln(&s)

	c := s[0]
	fmt.Println(int(c))
}
