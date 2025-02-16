package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	var err error
	var n, b, m float64
	var count int

loop:
	for {
		s, err = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		switch err {
		case nil:
		case io.EOF:
			break loop
		}

		n, _ = strconv.ParseFloat(strings.Split(s, " ")[0], 64)
		b, _ = strconv.ParseFloat(strings.Split(s, " ")[1], 64)
		m, _ = strconv.ParseFloat(strings.Split(s, " ")[2], 64)

		count = 0
		for n < m {
			n += (n * (b / 100))
			count++
		}
		fmt.Fprintln(writer, count)
	}

	writer.Flush()
}
