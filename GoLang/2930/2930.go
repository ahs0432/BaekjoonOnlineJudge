package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(b, r byte) int {
	if b == 'S' {
		if r == 'S' {
			return 1
		}
		if r == 'P' {
			return 2
		}
		if r == 'R' {
			return 0
		}
	} else if b == 'P' {
		if r == 'S' {
			return 0
		}
		if r == 'P' {
			return 1
		}
		if r == 'R' {
			return 2
		}
	} else if b == 'R' {
		if r == 'S' {
			return 2
		}
		if r == 'P' {
			return 0
		}
		if r == 'R' {
			return 1
		}
	}
	return 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var r, n int
	fmt.Fscanln(reader, &r)

	var s string
	fmt.Fscanln(reader, &s)

	fmt.Fscanln(reader, &n)
	f := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &f[i])
	}

	score := 0
	maxS := 0

	var win []int
	for i := 0; i < r; i++ {
		win = make([]int, 3)
		for j := 0; j < n; j++ {
			win[0] += check('S', f[j][i])
			win[1] += check('P', f[j][i])
			win[2] += check('R', f[j][i])
			score += check(s[i], f[j][i])
		}
		maxS += max(win[0], max(win[1], win[2]))
	}

	fmt.Println(score)
	fmt.Println(maxS)
}
