package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		tmp, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		if tmp == "#\n" {
			break
		} else {
			count := 0
			for _, i := range tmp {
				if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' || i == 'A' || i == 'E' || i == 'I' || i == 'O' || i == 'U' {
					count++
				}
			}
			fmt.Fprintln(writer, count)
		}
	}
	writer.Flush()
}
