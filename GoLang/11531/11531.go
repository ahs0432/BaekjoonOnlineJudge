package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var m int
	var name, res string

	count := 0
	ans := 0
	nameCount := map[string]int{}
	for {
		fmt.Fscanln(reader, &m, &name, &res)

		if m == -1 {
			break
		}

		if res == "right" {
			count += nameCount[name]*20 + m
			ans++
		}
		nameCount[name]++
	}

	fmt.Println(ans, count)
}
