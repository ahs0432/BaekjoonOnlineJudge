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

	var a, b []byte
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	for i := 0; i < n; i++ {
		for j := 0; j < len(a); j++ {
			if a[j] == '0' {
				a[j] = '1'
			} else {
				a[j] = '0'
			}
		}
	}

	if string(a) == string(b) {
		fmt.Println("Deletion succeeded")
	} else {
		fmt.Println("Deletion failed")
	}
}
