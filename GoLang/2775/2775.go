package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var k, n int
		fmt.Fscanln(reader, &k)
		fmt.Fscanln(reader, &n)

		table := make([][]int, k+1)

		for j := 0; j <= k; j++ {
			table[j] = make([]int, n)

			for l := 0; l < n; l++ {
				if j == 0 {
					table[j][l] = (l + 1)
				} else if l == 0 {
					table[j][l] = 1
				} else {
					table[j][l] = (table[j][l-1] + table[j-1][l])
				}
			}
		}

		fmt.Println(table[k][n-1])
	}
}

/*
	1	2	3	4	5	6	7	8	9
	1	3	6	10	15	21	28	36	45
	1	4	10	20	35	56	84	120	165
	1	5	15	35	70	126	210	330	495
	1	6	21	56	126	252	462	792	1287
*/
