package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s []byte
	fmt.Fscanln(reader, &s)

	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(string(s))
}
