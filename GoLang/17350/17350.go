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

	var tmp string
	t := false
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp == "anj" {
			t = true
		}
	}

	if t {
		fmt.Println("뭐야;")
	} else {
		fmt.Println("뭐야?")
	}
}
