package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		card1 := make([]string, n)
		card2 := make([]string, n)

		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &card1[j])
			} else {
				fmt.Fscan(reader, &card1[j])
			}
		}

		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &card2[j])
			} else {
				fmt.Fscan(reader, &card2[j])
			}
		}

		sort.Strings(card1)
		sort.Strings(card2)

		isCheater := false
		for j := 0; j < n; j++ {
			if card1[j] != card2[j] {
				isCheater = true
				break
			}
		}

		if isCheater {
			fmt.Fprintln(writer, "CHEATER")
		} else {
			fmt.Fprintln(writer, "NOT CHEATER")
		}
	}
	writer.Flush()
}
