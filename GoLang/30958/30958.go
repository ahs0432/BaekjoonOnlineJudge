package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	t := make([]int, 122-97)
	for _, c := range s {
		if c >= 97 && c <= 122 {
			t[c-97] += 1
		}
	}

	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})
	fmt.Println(t[0])
}
