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

	count := 0.0
	sum := 0.0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			sum += 4.0
		} else if s[i] == 'B' {
			sum += 3.0
		} else if s[i] == 'C' {
			sum += 2.0
		} else if s[i] == 'D' {
			sum += 1.0
		}

		if len(s) > i+1 && s[i+1] == '+' {
			sum += 0.5
			i++
		}
		count++
	}
	fmt.Printf("%.05f\n", sum/count)
}
