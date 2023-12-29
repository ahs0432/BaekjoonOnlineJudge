package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s1, s2 int
	fmt.Fscanln(reader, &s1, &s2)

	var a, b int
	for i := 0; i < s1; i++ {
		fmt.Fscanln(reader, &a, &b)

		if a != b {
			fmt.Println("Wrong Answer")
			return
		}
	}

	for i := 0; i < s2; i++ {
		fmt.Fscanln(reader, &a, &b)

		if a != b {
			fmt.Println("Why Wrong!!!")
			return
		}
	}

	fmt.Println("Accepted")
}
