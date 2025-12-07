package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var t int
	var song, str string
	table := make(map[string]string)
	tmp := make([]string, 7)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &t, &song, &tmp[0], &tmp[1], &tmp[2], &tmp[3], &tmp[4], &tmp[5], &tmp[6])
		str = tmp[0] + tmp[1] + tmp[2]

		if _, check := table[str]; !check {
			table[str] = song
		} else {
			table[str] = "?"
		}
	}

	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &tmp[0], &tmp[1], &tmp[2])
		str = tmp[0] + tmp[1] + tmp[2]

		if _, check := table[str]; !check {
			fmt.Println("!")
		} else {
			fmt.Println(table[str])
		}
	}
}
