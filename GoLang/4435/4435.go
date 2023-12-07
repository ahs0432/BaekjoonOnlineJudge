package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	g := make([]int, 6)
	s := make([]int, 7)

	var gs, ss int
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &g[0], &g[1], &g[2], &g[3], &g[4], &g[5])
		fmt.Fscanln(reader, &s[0], &s[1], &s[2], &s[3], &s[4], &s[5], &s[6])

		gs = g[0] + g[1]*2 + g[2]*3 + g[3]*3 + g[4]*4 + g[5]*10
		ss = s[0] + s[1]*2 + s[2]*2 + s[3]*2 + s[4]*3 + s[5]*5 + s[6]*10

		if gs > ss {
			fmt.Fprintf(writer, "Battle %d: Good triumphs over Evil\n", i)
		} else if gs < ss {
			fmt.Fprintf(writer, "Battle %d: Evil eradicates all trace of Good\n", i)
		} else {
			fmt.Fprintf(writer, "Battle %d: No victor on this battle field\n", i)
		}
	}

	writer.Flush()
}
