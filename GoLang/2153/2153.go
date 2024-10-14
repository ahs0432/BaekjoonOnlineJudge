package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	sum := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			sum += int(c) - 96
		} else {
			sum += int(c) - 38
		}
	}

	for i := 2; i <= sum/2; i++ {
		if sum%i == 0 {
			fmt.Println("It is not a prime word.")
			return
		}
	}

	fmt.Println("It is a prime word.")
}
