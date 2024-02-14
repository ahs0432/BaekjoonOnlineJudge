package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &s)

	ans := []byte{}
	for i := 0; i < len(s); i += n {
		ans = append(ans, s[i])
	}
	fmt.Println(string(ans))
}
