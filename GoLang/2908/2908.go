package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverse(str string) string {
	var tmp string

	for _, s := range str {
		tmp = string(s) + tmp
	}

	return tmp
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n1, n2 string
	fmt.Fscanln(reader, &n1, &n2)

	if len(n1) > len(n2) {
		fmt.Println(reverse(n1))
	} else if len(n1) < len(n2) {
		fmt.Println(reverse(n2))
	} else {
		n1Int, err := strconv.Atoi(reverse(n1))

		if err != nil {
			return
		}

		n2Int, err := strconv.Atoi(reverse(n2))

		if err != nil {
			return
		}

		if n1Int > n2Int {
			fmt.Println(n1Int)
		} else {
			fmt.Println(n2Int)
		}
	}
}
