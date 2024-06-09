package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	var s string
	for {
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}

		table := make([]int, 5)
		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &s)
			if s == "M" || s == "L" { // Joe
				table[0]++
			} else if s == "S" { // James
				table[3]++
			} else if s == "X" { // Dummy
				table[4]++
			} else {
				tmp := 0
				for j := 0; j < len(s); j++ {
					tmp *= 10
					tmp += int(s[j]) - 48
				}

				if tmp >= 12 { // Jean
					table[1]++
				} else { // Jane
					table[2]++
				}
			}
		}
		fmt.Fprintln(writer, table[0], table[1], table[2], table[3], table[4])
	}
	writer.Flush()
}
