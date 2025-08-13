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

	var spi []string
	var n, w, d, t, tmp int
	var s string
	var err error

loop:
	for {
		s, err = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		switch err {
		case nil:
		case io.EOF:
			break loop
		}

		spi = strings.Split(s, " ")
		n, _ = strconv.Atoi(spi[0])
		w, _ = strconv.Atoi(spi[1])
		d, _ = strconv.Atoi(spi[2])
		t, _ = strconv.Atoi(spi[3])

		tmp = ((w * ((n - 1) * n) / 2) - t) / d
		if tmp == 0 {
			fmt.Fprintln(writer, n)
		} else {
			fmt.Fprintln(writer, tmp)
		}
	}
	writer.Flush()
}
