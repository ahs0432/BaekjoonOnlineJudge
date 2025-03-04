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

	var s string
	var w []string
	var now byte

	check := true
	for {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		if s == "*" {
			break
		}

		s = strings.ToUpper(s)
		w = strings.Split(s, " ")
		now = w[0][0]
		check = true

		for j := 1; j < len(w); j++ {
			if now != w[j][0] {
				check = false
				break
			}
		}

		if check {
			fmt.Fprintln(writer, "Y")
		} else {
			fmt.Fprintln(writer, "N")
		}

	}
	writer.Flush()
}
