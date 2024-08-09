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

	var n string
	fmt.Fscanln(reader, &n)

	oct := []string{"000", "001", "010", "011", "100", "101", "110", "111"}

	for i := 0; i < len(n); i++ {
		if i == 0 {
			tmp, _ := strconv.Atoi(oct[n[i]-'0'])
			fmt.Fprint(writer, tmp)
			continue
		}
		fmt.Fprint(writer, oct[n[i]-'0'])
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
