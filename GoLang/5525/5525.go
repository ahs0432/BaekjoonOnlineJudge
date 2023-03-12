package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	var s string
	fmt.Fscanln(reader, &s)

	n = n*2 + 1

	table := make([]int, 0)
	count := 0
	for i := 0; i < len(s); i++ {
		if count%2 == 0 && s[i] == 'I' {
			count++
		} else if count%2 == 1 && s[i] == 'O' {
			count++
		} else {
			if count > 2 {
				if count%2 == 0 {
					count -= 1
				}

				table = append(table, count)
			}

			if s[i] == 'I' {
				count = 1
			} else {
				count = 0
			}
		}
	}

	if count > 2 {
		if count%2 == 0 {
			count -= 1
		}

		table = append(table, count)
	}

	check := 0
	for _, v := range table {
		tmp := v - n
		if tmp >= 0 {
			check++
			check += (tmp / 2)
		}
	}
	fmt.Println(check)
}
