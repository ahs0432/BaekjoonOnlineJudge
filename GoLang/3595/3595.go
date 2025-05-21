package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscanln(reader, &n)

	var a, b, c int64
	var sum, min int64 = 0, 2147483647
	for i := int64(1); i <= n; i++ {
		for j := int64(1); j <= i; j++ {
			if i*j > n {
				break
			}

			for k := int64(1); k <= j; k++ {
				if i*j*k > n {
					break
				}

				if i*j*k == n {
					sum = i*j + j*k + k*i
					if sum < min {
						min = sum
						a = i
						b = j
						c = k
					}
				}
			}
		}
	}

	fmt.Println(a, b, c)
}
