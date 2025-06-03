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
	var n, now int64
	var checkNum, sum int
	var check []bool
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

		n, _ = strconv.ParseInt(s, 10, 64)

		checkNum = 10
		check = make([]bool, 10)

		now, sum = n, 1
		for {
			s = fmt.Sprintf("%d", now)

			for i := 0; i < len(s); i++ {
				if !check[s[i]-'0'] {
					check[s[i]-'0'] = true
					checkNum--
				}
			}

			if checkNum == 0 {
				break
			}

			sum++
			now += n
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
