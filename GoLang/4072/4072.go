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

	var tmp string
	var strs []string
	for {
		tmp, _ = reader.ReadString('\n')
		tmp = strings.TrimSuffix(tmp, "\n")

		if tmp == "#" {
			break
		}

		strs = strings.Split(tmp, " ")

		for i, str := range strs {
			for i := len(str) - 1; i >= 0; i-- {
				fmt.Fprint(writer, string(str[i]))
			}

			if i != len(strs)-1 {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
