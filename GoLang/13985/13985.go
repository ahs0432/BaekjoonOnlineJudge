package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, _ := reader.ReadString('\n')
	sum := (int(n[0]) - 48) + (int(n[4]) - 48)

	if (int(n[8]) - 48) == sum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
