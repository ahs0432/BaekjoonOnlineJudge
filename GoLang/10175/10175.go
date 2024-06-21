package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, now int
	fmt.Fscanln(reader, &n)

	var location, species string
	animal := []string{"Bobcat", "Coyote", "Mountain Lion", "Wolf"}

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &location, &species)

		solo, best, score := false, 0, 0
		table := make([]int, 4)

		for _, c := range species {
			if c == 'B' {
				now = 0
				table[now] += 2
			} else if c == 'C' {
				now = 1
				table[now] += 1
			} else if c == 'M' {
				now = 2
				table[now] += 4
			} else if c == 'W' {
				now = 3
				table[now] += 3
			}

			if table[now] > score {
				solo = true
				best = now
				score = table[now]
			} else if table[now] == score {
				solo = false
			}
		}

		if solo {
			fmt.Fprintf(writer, "%s: The %s is the dominant species\n", location, animal[best])
		} else {
			fmt.Fprintf(writer, "%s: There is no dominant species\n", location)
		}
	}
	writer.Flush()
}
