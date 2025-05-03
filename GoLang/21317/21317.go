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

	smallJump := make([]int, n-1)
	bigJump := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscanln(reader, &smallJump[i], &bigJump[i])
	}

	var k int
	fmt.Fscanln(reader, &k)

	rocks := [][]int{make([]int, n), make([]int, n)}

	for i := 0; i < n; i++ {
		rocks[0][i] = 100001
		rocks[1][i] = 100001
	}

	rocks[0][0] = 0
	if n > 1 {
		rocks[0][1] = smallJump[0]
	}
	if n > 2 {
		rocks[0][2] = min(smallJump[0]+smallJump[1], bigJump[0])
	}
	for i := 3; i < n; i++ {
		rocks[0][i] = min(smallJump[i-1]+rocks[0][i-1], bigJump[i-2]+rocks[0][i-2])
		rocks[1][i] = min(smallJump[i-1]+rocks[1][i-1], bigJump[i-2]+rocks[1][i-2], k+rocks[0][i-3])
	}

	fmt.Println(min(rocks[0][n-1], rocks[1][n-1]))
}
