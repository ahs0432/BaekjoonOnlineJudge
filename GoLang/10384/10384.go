package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	var err error

	var min int
	var alpha []int
loop:
	for i := 1; i <= n; i++ {
		s, err = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		switch err {
		case nil:
		case io.EOF:
			break loop
		}

		alpha = make([]int, 26)
		for j := 0; j < len(s); j++ {
			if s[j] >= 'a' && s[j] <= 'z' {
				alpha[s[j]-'a']++
			} else if s[j] >= 'A' && s[j] <= 'Z' {
				alpha[s[j]-'A']++
			}
		}

		min = 2147483647
		for j := 0; j < 26; j++ {
			if alpha[j] < min {
				min = alpha[j]
			}
		}

		fmt.Fprintf(writer, "Case %d: ", i)
		switch min {
		case 0:
			fmt.Fprintln(writer, "Not a pangram")
		case 1:
			fmt.Fprintln(writer, "Pangram!")
		case 2:
			fmt.Fprintln(writer, "Double pangram!!")
		case 3:
			fmt.Fprintln(writer, "Triple pangram!!!")
		}
	}
	writer.Flush()
}
