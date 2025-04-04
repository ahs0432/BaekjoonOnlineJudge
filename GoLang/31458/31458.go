package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var now int
	var tmp string
	var s []string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if strings.Contains(tmp, "1") {
			s = strings.Split(tmp, "1")
			now = 1
		} else {
			s = strings.Split(tmp, "0")
			now = 0
		}

		if len(s[1]) >= 1 {
			now = 1
		}

		if len(s[0])%2 == 1 {
			if now == 1 {
				now = 0
			} else {
				now = 1
			}
		}
		fmt.Fprintln(writer, now)
	}
	writer.Flush()
}
