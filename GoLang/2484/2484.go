package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp, sum int
	for i := 0; i < n; i++ {
		table := make([]int, 7)

		for j := 0; j < 4; j++ {
			fmt.Fscan(reader, &tmp)
			table[tmp]++
		}

		for j := 6; j > 0; j-- {
			if table[j] == 4 {
				if sum < 50000+5000*j {
					sum = 50000 + 5000*j
				}
				break
			} else if table[j] == 3 {
				if sum < 10000+1000*j {
					sum = 10000 + 1000*j
				}
				break
			} else if table[j] == 2 {
				for k := j - 1; k > 0; k-- {
					if table[k] == 2 {
						if sum < 2000+j*500+k*500 {
							sum = 2000 + j*500 + k*500
						}
						break
					} else {
						if sum < 1000+100*j {
							sum = 1000 + 100*j
						}
					}
				}
				break
			} else if table[j] == 1 {
				if sum < 100*j {
					sum = 100 * j
				}
			}
		}
	}
	fmt.Println(sum)
}
