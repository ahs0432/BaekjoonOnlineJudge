package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Team struct {
	ID    int
	Score int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	teams := make([]Team, n)
	for i := 0; i < n; i++ {
		teams[i].ID = i + 1
	}

	var a, b, c, d int
	for {
		_, err := fmt.Fscanln(reader, &a, &b, &c, &d)
		if err != nil {
			break
		}
		a--
		b--

		if c < d {
			teams[b].Score += 3
		} else if c > d {
			teams[a].Score += 3
		} else {
			teams[a].Score++
			teams[b].Score++
		}
	}

	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Score > teams[j].Score
	})

	ans := make([]int, n+1)
	rank := 1
	ans[teams[0].ID] = rank

	for i := 1; i < n; i++ {
		if teams[i].Score != teams[i-1].Score {
			rank = i + 1
		}
		ans[teams[i].ID] = rank
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(writer, ans[i])
	}
}
