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

	var s string
	fmt.Fscanln(reader, &s)

	table := strings.Split(s, ",")

	var n int
	sum := 0
	for i := 0; i < len(table); i++ {
		n, _ = strconv.Atoi(table[i])
		sum += n
	}
	fmt.Println(sum)
}
