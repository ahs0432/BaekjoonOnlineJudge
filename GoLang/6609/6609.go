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

	var tmp []string
	var m, p, l, e, r, s, n, sum int
loop:
	for {
		a, err := reader.ReadString('\n')
		a = strings.TrimSuffix(a, "\n")
		switch err {
		case nil:
		case io.EOF:
			break loop
		}

		if a == "" {
			break
		}

		tmp = strings.Split(a, " ")
		m, _ = strconv.Atoi(tmp[0])
		p, _ = strconv.Atoi(tmp[1])
		l, _ = strconv.Atoi(tmp[2])
		e, _ = strconv.Atoi(tmp[3])
		r, _ = strconv.Atoi(tmp[4])
		s, _ = strconv.Atoi(tmp[5])
		n, _ = strconv.Atoi(tmp[6])

		for i := 0; i < n; i++ {
			sum = m
			m = p / s
			p = l / r
			l = sum * e
		}
		fmt.Fprintln(writer, m)
	}
	writer.Flush()
}
