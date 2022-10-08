package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var n string
		fmt.Fscanln(reader, &n)

		if n == "0" {
			break
		} else if len(n) == 1 {
			fmt.Println("yes")
		} else {
			check := true
			for i := 0; i < len(n)/2; i++ {
				if n[i] != n[len(n)-1-i] {
					check = false
					break
				}
			}

			if check {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}
