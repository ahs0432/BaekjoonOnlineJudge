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

	var h, p float64

loop:
	for {
		s, err := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		switch err {
		case nil:
		case io.EOF:
			break loop
		}
		sp := strings.Split(s, " ")
		h, _ = strconv.ParseFloat(sp[0], 64)
		p, _ = strconv.ParseFloat(sp[1], 64)
		fmt.Fprintf(writer, "%.2f\n", h/p)
	}
	writer.Flush()
}
