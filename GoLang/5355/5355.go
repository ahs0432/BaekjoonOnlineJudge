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
	var f float64
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		slices := strings.Split(s, " ")
		f, _ = strconv.ParseFloat(slices[0], 64)

		for j := 1; j < len(slices); j++ {
			if slices[j] == "@" {
				f *= 3
			} else if slices[j] == "%" {
				f += 5
			} else if slices[j] == "#" {
				f -= 7
			}
		}

		fmt.Fprintf(writer, "%.2f\n", f)
	}

	writer.Flush()
}
