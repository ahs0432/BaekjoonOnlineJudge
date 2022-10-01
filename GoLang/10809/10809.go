package main

import "fmt"

func main() {
	alpha := make([]int, 26)
	for i := 0; i < len(alpha); i++ {
		alpha[i] = -1
	}

	var s string
	fmt.Scanln(&s)

	for i := 0; i < len(s); i++ {
		if alpha[int(s[i])-97] == -1 {
			alpha[int(s[i])-97] = i
		}
	}

	for i := 0; i < len(alpha); i++ {
		fmt.Print(alpha[i], " ")
	}
}
