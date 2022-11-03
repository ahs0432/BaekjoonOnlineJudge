package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	set := make([]bool, 20)

	for i := 0; i < n; i++ {
		var word string
		var m int

		fmt.Fscanln(reader, &word, &m)

		if word == "add" {
			set[m-1] = true
		} else if word == "remove" {
			set[m-1] = false
		} else if word == "check" {
			if set[m-1] {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		} else if word == "toggle" {
			set[m-1] = !set[m-1]
		} else if word == "all" {
			for j := 0; j < 20; j++ {
				set[j] = true
			}
		} else {
			for j := 0; j < 20; j++ {
				set[j] = false
			}
		}
	}

	writer.Flush()
}
