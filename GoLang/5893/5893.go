package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscanln(reader, &n)

	res := append([]byte(n), []byte{'0', '0', '0', '0'}...)
	respos := len(res)

	for i := len(n) - 1; i >= 0; i-- {
		respos--
		if n[i] == '0' {
			continue
		}
		res[respos]++
	}

	for i := len(res) - 1; i > 0; i-- {
		if res[i] >= '2' {
			res[i-1] += (res[i] - '0') / 2
			res[i] = ((res[i] - '0') % 2) + '0'
		}
	}

	if res[0] == '2' {
		res[0] = '0'
		fmt.Print("1")
	} else if res[0] == '3' {
		res[0] = '1'
		fmt.Print("1")
	} else if res[0] == '4' {
		res[0] = '0'
		fmt.Print("10")
	}

	fmt.Println(string(res))
}
