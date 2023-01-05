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

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)
		if len(tmp) >= 6 && len(tmp) <= 9 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
