package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp1, tmp2 string
	table := map[string]string{}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp1, &tmp2)
		table[tmp2] = tmp1
	}

	var tmp string
	fmt.Fscanln(reader, &tmp)

	var merge string
	for i := 0; i < len(tmp); i++ {
		merge += table[string(tmp[i])]
	}

	var s, e int
	fmt.Fscanln(reader, &s, &e)
	fmt.Println(merge[s-1 : e])
}
