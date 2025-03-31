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
		var d1, d2 int
		fmt.Fscanln(reader, &d1, &d2)
		fmt.Print("Case ")
		fmt.Print(i + 1)
		fmt.Println(":", d1+d2)
	}
}
