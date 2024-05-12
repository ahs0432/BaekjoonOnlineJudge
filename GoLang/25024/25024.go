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

	var x, y int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x, &y)

		if x <= 23 && y <= 59 {
			fmt.Fprint(writer, "Yes ")
		} else {
			fmt.Fprint(writer, "No ")
		}

		if x >= 1 && x <= 12 {
			if ((x == 1 || x == 3 || x == 5 || x == 7 || x == 8 || x == 10 || x == 12) && (y >= 1 && y <= 31)) ||
				((x == 4 || x == 6 || x == 9 || x == 11) && (y >= 1 && y <= 30)) ||
				((x == 2) && (y >= 1 && y <= 29)) {
				fmt.Fprint(writer, "Yes")
			} else {
				fmt.Fprint(writer, "No")
			}
		} else {
			fmt.Fprint(writer, "No")
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
