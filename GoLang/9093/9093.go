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

	for i := 0; i < n; i++ {
		s, err := reader.ReadString('\n')
		if nil != err {
			break
		}

		s = strings.TrimSuffix(s, "\n")
		sp := strings.Split(s, " ")

		for _, p := range sp {
			for j := len(p) - 1; j >= 0; j-- {
				fmt.Fprint(writer, string(p[j]))
			}
			fmt.Fprint(writer, " ")
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
