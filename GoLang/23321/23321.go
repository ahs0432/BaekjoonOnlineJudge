package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	table := make([]string, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	photos := make([][]rune, 5)
	var tmp string
	for i := 0; i < len(table[0]); i++ {
		tmp = string([]byte{
			table[0][i],
			table[1][i],
			table[2][i],
			table[3][i],
			table[4][i],
		})

		if tmp == ".omln" {
			photos[0] = append(photos[0], 'o')
			photos[1] = append(photos[1], 'w')
			photos[2] = append(photos[2], 'l')
			photos[3] = append(photos[3], 'n')
			photos[4] = append(photos[4], '.')
		} else if tmp == "owln." {
			photos[0] = append(photos[0], '.')
			photos[1] = append(photos[1], 'o')
			photos[2] = append(photos[2], 'm')
			photos[3] = append(photos[3], 'l')
			photos[4] = append(photos[4], 'n')
		} else {
			photos[0] = append(photos[0], '.')
			photos[1] = append(photos[1], '.')
			photos[2] = append(photos[2], 'o')
			photos[3] = append(photos[3], 'L')
			photos[4] = append(photos[4], 'n')
		}
	}

	for _, photo := range photos {
		fmt.Fprintln(writer, string(photo))
	}
	writer.Flush()
}
