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

	for {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		if s == "END" {
			break
		} else {
			tmp := ""
			for _, c := range s {
				tmp = string(c) + tmp
			}
			fmt.Fprintln(writer, tmp)
		}
	}
	writer.Flush()
}
