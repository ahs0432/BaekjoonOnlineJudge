package main

import (
	"bufio"
	"fmt"
	"os"
)

func search(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for k := 0; k < len(arr); k++ {
				if i != j && i != k && j != k && (arr[i]-arr[j])%arr[k] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i])
	}
	if search(slice) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
