package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func choice(i int, t bool) string {
	if i == 1 {
		if t {
			return "Habb Yakk"
		}
		return "Yakk"
	} else if i == 2 {
		if t {
			return "Dobara"
		}
		return "Doh"
	} else if i == 3 {
		if t {
			return "Dousa"
		}
		return "Seh"
	} else if i == 4 {
		if t {
			return "Dorgy"
		}
		return "Ghar"
	} else if i == 5 {
		if t {
			return "Dabash"
		}
		return "Bang"
	} else {
		if t {
			return "Dosh"
		}
		return "Sheesh"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]int, 2)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &table[0], &table[1])
		sort.Slice(table, func(j, k int) bool {
			return table[j] > table[k]
		})

		if table[0] == table[1] {
			fmt.Fprintf(writer, "Case %d: %s\n", i, choice(table[0], true))
		} else {
			if table[0] == 6 && table[1] == 5 {
				fmt.Fprintf(writer, "Case %d: %s %s\n", i, "Sheesh", "Beesh")
			} else {
				fmt.Fprintf(writer, "Case %d: %s %s\n", i, choice(table[0], false), choice(table[1], false))
			}
		}
	}
	writer.Flush()
}
