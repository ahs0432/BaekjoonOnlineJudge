package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, sum int
	fmt.Fscanln(reader, &n)

	var tmp string
	for i := 0; i < n; i++ {
		tmp, _ = reader.ReadString('\n')
		tmp = strings.TrimSuffix(tmp, "\n")

		sum = 0
		ss := strings.Split(tmp, " ")
		for _, s := range ss {
			t, _ := strconv.Atoi(s)
			sum += t
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
