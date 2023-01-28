package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var p, x string
		var n int
		fmt.Fscanln(reader, &p)
		fmt.Fscanln(reader, &n)
		fmt.Fscanln(reader, &x)

		st := true
		first := 0
		last := n

		for _, pp := range p {
			if string(pp) == "R" {
				st = !st
			} else {
				if st {
					first += 1
				} else {
					last -= 1
				}
			}
		}

		if n < first || last < first {
			fmt.Fprintln(writer, "error")
			writer.Flush()
			continue
		}

		table := make([]int, n)
		tmp := strings.Split(strings.Split(strings.Split(x, "[")[1], "]")[0], ",")

		if len(tmp) != 1 || tmp[0] != "" {
			for j, num := range tmp {
				table[j], _ = strconv.Atoi(num)
			}
		}

		fmt.Fprint(writer, "[")
		if st {
			for j := first; j < last; j++ {
				fmt.Fprint(writer, table[j])
				if j != (last - 1) {
					fmt.Fprint(writer, ",")
				}
			}
		} else {
			for j := last - 1; j >= first; j-- {
				fmt.Fprint(writer, table[j])
				if j != first {
					fmt.Fprint(writer, ",")
				}
			}
		}
		fmt.Fprintln(writer, "]")
		writer.Flush()
	}
}
