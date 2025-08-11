package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s1, s2 string
	var score1, score2 int

	for {
		fmt.Fscanln(reader, &s1)
		fmt.Fscanln(reader, &s2)

		if s1 == "E" {
			break
		}

		score1, score2 = 0, 0
		for i := range s1 {
			if s1[i] == s2[i] {
				continue
			} else if s1[i] == 'R' && s2[i] == 'S' {
				score1++
			} else if s1[i] == 'S' && s2[i] == 'P' {
				score1++
			} else if s1[i] == 'P' && s2[i] == 'R' {
				score1++
			} else {
				score2++
			}
		}

		fmt.Fprintln(writer, "P1:", score1)
		fmt.Fprintln(writer, "P2:", score2)
	}
	writer.Flush()
}
