package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func selSort(a, b []int) int {
	if reflect.DeepEqual(a, b) {
		return 1
	}

	var max int
	for i := len(a) - 1; i > 0; i-- {
		max = 0
		for j := 0; j <= i; j++ {
			if a[j] > a[max] {
				max = j
			}
		}

		if max != i {
			a[i], a[max] = a[max], a[i]
		}

		if reflect.DeepEqual(a, b) {
			return 1
		}
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &a[i])
		} else {
			fmt.Fscan(reader, &a[i])
		}
	}

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &b[i])
		} else {
			fmt.Fscan(reader, &b[i])
		}
	}
	fmt.Println(selSort(a, b))
}
