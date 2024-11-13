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

	fruits := map[string]int{}

	var s string
	var x int

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s, &x)
		fruits[s] += x
	}

	ans := "NO"
	for _, v := range fruits {
		if v == 5 {
			ans = "YES"
		}
	}
	fmt.Println(ans)
}
