package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	check := -1
	count := 0
	var tmp int
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				tmp = a[j+1]
				a[j+1] = a[j]
				a[j] = tmp
				count++
			}
			if count == k {
				check = j
				break
			}
		}

		if check != -1 {
			break
		}
	}

	if check == -1 {
		fmt.Println(check)
	} else {
		fmt.Println(a[check], a[check+1])
	}
}
