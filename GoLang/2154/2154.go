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

	var n int
	fmt.Fscanln(reader, &n)

	var s strings.Builder
	for i := 1; i <= n; i++ {
		s.WriteString(strconv.Itoa(i))
	}

	fmt.Println(strings.Index(s.String(), strconv.Itoa(n)) + 1)
}
