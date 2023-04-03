package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	if m < 3 {
		fmt.Println("NEWBIE!")
	} else if m <= n {
		fmt.Println("OLDBIE!")
	} else {
		fmt.Println("TLE!")
	}
}
