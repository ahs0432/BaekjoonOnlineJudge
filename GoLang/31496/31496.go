package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n, &s)

	count := 0
	var tmpItem string
	var tmpAmount int
	var tmpTable []string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmpItem, &tmpAmount)
		tmpTable = strings.Split(tmpItem, "_")

		for _, c := range tmpTable {
			if c == s {
				count += tmpAmount
				break
			}
		}
	}
	fmt.Println(count)
}
