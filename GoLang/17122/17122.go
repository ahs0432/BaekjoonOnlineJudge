package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	table := []bool{
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
		true, false, true, false, true, false, true, false,
		false, true, false, true, false, true, false, true,
	}

	var n int
	fmt.Fscanln(reader, &n)

	var str string
	var t, tmp1 int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &str, &t)
		tmp1 = 8*int(str[1]-'1') + int(str[0]-'A'+1)

		if table[tmp1-1] == table[t-1] {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
	writer.Flush()
}
