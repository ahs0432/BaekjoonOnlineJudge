package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	bus := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &bus[i])
	}

	arr := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			fmt.Fscan(reader, &arr[i][j])
		}
	}

	sum := 0
	for i := 1; i < m; i++ {
		sum += arr[bus[i-1]][bus[i]]
	}
	fmt.Println(sum)
}
