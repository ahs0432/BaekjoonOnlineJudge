package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	table := []string{"1QAZ", "2WSX", "3EDC", "4RFV5TGB", "6YHN7UJM", "8IK,", "9OL.", "0P;/-['=]"}
	ans := make([]int, 8)

	fmt.Fscanln(reader, &s)
	for _, c := range s {
		for i := 0; i < 8; i++ {
			if strings.Contains(table[i], string(c)) {
				ans[i]++
			}
		}
	}

	for i := 0; i < 8; i++ {
		fmt.Fprintln(writer, ans[i])
	}
	writer.Flush()
}
