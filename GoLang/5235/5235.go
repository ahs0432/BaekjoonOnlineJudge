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

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	var odd, even int
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		odd, even = 0, 0
		s = strings.TrimSpace(s)
		sc := strings.Split(s, " ")[1:]

		for _, c := range sc {
			it, _ := strconv.Atoi(c)
			if it%2 == 0 {
				even += it
			} else {
				odd += it
			}
		}

		if odd == even {
			fmt.Fprintln(writer, "TIE")
		} else if odd > even {
			fmt.Fprintln(writer, "ODD")
		} else {
			fmt.Fprintln(writer, "EVEN")
		}
	}
	writer.Flush()
}
