package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscanln(reader, &n)

	if n <= 32767 && n >= -32768 {
		fmt.Println("short")
	} else if n <= 2147483647 && n >= -2147483648 {
		fmt.Println("int")
	} else {
		fmt.Println("long long")
	}
}
