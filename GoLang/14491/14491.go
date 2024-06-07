package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	ans := ""
	for n != 0 {
		ans += strconv.Itoa(n % 9)
		n /= 9
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(writer, string(ans[i]))
	}

	fmt.Fprintln(writer)
	writer.Flush()
}
