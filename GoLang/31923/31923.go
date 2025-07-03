package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, p, q int
	fmt.Fscanln(reader, &n, &p, &q)

	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	var check bool
	if p == q {
		check = true
		for i := 0; i < n; i++ {
			if a[i] != b[i] {
				check = false
				break
			}
		}

		if check {
			fmt.Println("YES")
			for i := 0; i < n; i++ {
				fmt.Print("0 ")
			}
			fmt.Println()
		} else {
			fmt.Println("NO")
		}
	} else {
		answer := make([]int, n)
		if p > q {
			p, q = q, p
			for i := 0; i < n; i++ {
				a[i], b[i] = b[i], a[i]
			}
		}

		check = true
		for i := 0; i < n; i++ {
			if a[i] == b[i] {
				answer[i] = 0
			} else if a[i] > b[i] {
				if (a[i]-b[i])%(q-p) == 0 {
					answer[i] = (a[i] - b[i]) / (q - p)
				} else {
					check = false
					break
				}
			} else {
				check = false
				break
			}
		}

		if check {
			fmt.Println("YES")
			for i := 0; i < n; i++ {
				fmt.Print(answer[i], " ")
			}
			fmt.Println()
		} else {
			fmt.Println("NO")
		}
	}
}
