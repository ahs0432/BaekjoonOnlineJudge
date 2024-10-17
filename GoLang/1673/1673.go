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

	var ans, n, k int

loop:
	for {
		s, err := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		switch err {
		case nil:
		case io.EOF:
			break loop
		}

		n, _ = strconv.Atoi(strings.Split(s, " ")[0])
		k, _ = strconv.Atoi(strings.Split(s, " ")[1])

		ans = 0
		ans += n

		for n/k > 0 {
			ans += n / k
			n = n/k + n%k
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
