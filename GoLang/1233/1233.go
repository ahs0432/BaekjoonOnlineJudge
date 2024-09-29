package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s1, s2, s3 int
	fmt.Fscanln(reader, &s1, &s2, &s3)

	table := make([]int, s1+s2+s3+1)

	c := 0
	max := 0
	for i := 1; i <= s1; i++ {
		for j := 1; j <= s2; j++ {
			for k := 1; k <= s3; k++ {
				table[i+j+k]++

				if table[i+j+k] > max {
					max = table[i+j+k]
					c = i + j + k
				} else if table[i+j+k] == max && c > i+j+k {
					c = i + j + k
				}
			}
		}
	}

	fmt.Println(c)
}
