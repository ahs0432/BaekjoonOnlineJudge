package main

import (
	"bufio"
	"fmt"
	"os"
)

func recursion(s string, l, r, count int) (int, int) {
	count++
	if l >= r {
		return 1, count
	} else if s[l] != s[r] {
		return 0, count
	} else {
		return recursion(s, l+1, r-1, count)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(reader, &t)

	s := make([]string, t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &s[i])
		fmt.Println(recursion(s[i], 0, len(s[i])-1, 0))
	}
}
