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
	writer := bufio.NewWriter(os.Stdout)

	var n int
	var tmp []string
	for {
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		}

		tmp = make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscanln(reader, &tmp[i])
		}

		sort.Slice(tmp, func(i, j int) bool {
			return strings.ToLower(tmp[i]) < strings.ToLower(tmp[j])
		})
		fmt.Fprintln(writer, tmp[0])
	}
	writer.Flush()
}
