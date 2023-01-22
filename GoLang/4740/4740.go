package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		n, _ := reader.ReadString('\n')
		n = n[:len(n)-1]

		if n == "***" {
			break
		} else {
			for i := len(n) - 1; i >= 0; i-- {
				fmt.Print(string(n[i]))
			}
			fmt.Println()
		}
	}

}
