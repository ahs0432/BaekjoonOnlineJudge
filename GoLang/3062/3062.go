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

	var t int
	fmt.Fscanln(reader, &t)

	var n string
	var tmp string
	var tmpr string
	var check bool
	var tmpi, tmpri int
	for i := 0; i < t; i++ {
		tmpr = ""
		fmt.Fscanln(reader, &tmp)

		for j := len(tmp) - 1; j >= 0; j-- {
			tmpr += string(tmp[j])
		}

		tmpi, _ = strconv.Atoi(tmp)
		tmpri, _ = strconv.Atoi(tmpr)
		n = strconv.Itoa(tmpi + tmpri)

		check = true
		for j := 0; j < len(n)/2; j++ {
			if n[j] != n[len(n)-1-j] {
				check = false
				break
			}
		}

		if check {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
	writer.Flush()
}
