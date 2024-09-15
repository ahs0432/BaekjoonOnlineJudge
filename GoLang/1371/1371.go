package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	max := 0
	table := make([]int, 26)
loop:
	for {
		s, err := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		s = strings.ReplaceAll(s, " ", "")

		switch err {
		case nil:
		case io.EOF:
			break loop
		}

		for _, c := range s {
			table[c-97]++
			if table[c-97] > max {
				max = table[c-97]
			}
		}
	}

	for i := 0; i < 26; i++ {
		if table[i] == max {
			fmt.Print(string(byte(i + 97)))
		}
	}
	fmt.Println()
}
