package main

import (
	"bufio"
	"fmt"
	"os"
)

func binToStr(s string) string {
	now := 1
	ans := 0
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == 'I' {
			ans += now
		}
		now *= 2
	}

	return string(byte(ans))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var b int
	var s, t string
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &b)
		fmt.Fscanln(reader, &s)

		t = ""
		for j := 0; j < b; j++ {
			t += binToStr(s[8*j : 8*(j+1)])
		}
		fmt.Fprintf(writer, "Case #%d: %s\n", i, t)
	}
	writer.Flush()
}
