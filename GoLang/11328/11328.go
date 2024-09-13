package main

import (
	"bufio"
	"fmt"
	"os"
)

func strfry(str1, str2 string) bool {
	tmp1 := make([]int, 26)
	tmp2 := make([]int, 26)

	for _, c := range str1 {
		tmp1[c-97]++
	}

	for _, c := range str2 {
		tmp2[c-97]++
	}

	for i := 0; i < 26; i++ {
		if tmp1[i] != tmp2[i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var str1, str2 string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &str1, &str2)
		if strfry(str1, str2) {
			fmt.Fprintln(writer, "Possible")
		} else {
			fmt.Fprintln(writer, "Impossible")
		}
	}
	writer.Flush()
}
