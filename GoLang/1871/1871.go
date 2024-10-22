package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var a, b, tmp int
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)

		a = (int(s[0])-65)*26*26 + (int(s[1])-65)*26 + (int(s[2]) - 65)
		b = (int(s[4])-48)*1000 + (int(s[5])-48)*100 + (int(s[6])-48)*10 + (int(s[7]) - 48)

		tmp = b - a

		if tmp < 0 {
			tmp *= -1
		}

		if tmp <= 100 {
			fmt.Fprintln(writer, "nice")
		} else {
			fmt.Fprintln(writer, "not nice")
		}
	}
	writer.Flush()
}
