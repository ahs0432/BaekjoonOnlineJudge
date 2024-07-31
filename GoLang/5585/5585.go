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

	t := 1000 - n
	ans := t / 500
	t %= 500
	ans += t / 100
	t %= 100
	ans += t / 50
	t %= 50
	ans += t / 10
	t %= 10
	ans += t / 5
	t %= 5
	ans += t / 1
	t %= 1
	fmt.Println(ans)
}
