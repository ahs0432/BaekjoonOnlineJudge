package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	countStreak := make(map[int]struct{})
	type Record struct {
		maxStreak int
		name      string
	}
	kpscList := make([]Record, n)

	var maxStreak, tmp int
	var name, result string
	for i := 0; i < n; i++ {
		maxStreak = 0
		tmp = 0
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &result)
			if result == "." {
				tmp++
			} else {
				tmp = 0
			}
			if tmp > maxStreak {
				maxStreak = tmp
			}
		}

		fmt.Fscanln(reader, &name)
		countStreak[maxStreak] = struct{}{}
		kpscList[i] = Record{maxStreak, name}
	}

	fmt.Fprintln(writer, len(countStreak))
	for _, k := range kpscList {
		fmt.Fprintf(writer, "%d %s\n", k.maxStreak, k.name)
	}
	writer.Flush()
}
