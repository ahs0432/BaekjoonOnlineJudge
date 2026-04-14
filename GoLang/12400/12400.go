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

	ans := "yhesocvxduiglbkrztnwjpfmaq"

	var t int
	fmt.Fscanln(reader, &t)

	for i := 1; i <= t; i++ {
		s, _ := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		fmt.Fprintf(writer, "Case #%d: ", i)
		for _, w := range s {
			if w >= 'a' && w <= 'z' {
				fmt.Fprintf(writer, "%c", ans[w-'a'])
			} else {
				fmt.Fprintf(writer, "%c", w)
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
