package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	var rank int
	results := make([]string, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &rank)
		p := (rank * 100) / n

		switch {
		case p >= 0 && p <= 4:
			results[i] = "1"
		case p <= 11:
			results[i] = "2"
		case p <= 23:
			results[i] = "3"
		case p <= 40:
			results[i] = "4"
		case p <= 60:
			results[i] = "5"
		case p <= 77:
			results[i] = "6"
		case p <= 89:
			results[i] = "7"
		case p <= 96:
			results[i] = "8"
		case p <= 100:
			results[i] = "9"
		default:
			results[i] = "0"
		}
	}

	fmt.Println(strings.Join(results, " "))
}
