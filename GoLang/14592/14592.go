package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type score struct {
	index int
	small int
	count int
	large int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	scores := make([]score, n)
	for i := 0; i < n; i++ {
		scores[i].index = i + 1
		fmt.Fscanln(reader, &scores[i].small, &scores[i].count, &scores[i].large)
	}

	sort.Slice(scores, func(i, j int) bool {
		if scores[i].small > scores[j].small {
			return true
		} else if scores[i].small == scores[j].small {
			if scores[i].count < scores[j].count {
				return true
			} else if scores[i].count == scores[j].count {
				if scores[i].large < scores[j].large {
					return true
				}
			}
		}
		return false
	})

	fmt.Println(scores[0].index)
}
