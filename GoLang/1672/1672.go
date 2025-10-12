package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(c byte) int {
	switch c {
	case 'A':
		return 0
	case 'G':
		return 1
	case 'C':
		return 2
	default:
		return 3
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := [][]byte{
		{'A', 'C', 'A', 'G'},
		{'C', 'G', 'T', 'A'},
		{'A', 'T', 'C', 'G'},
		{'G', 'A', 'G', 'T'},
	}

	var n int
	fmt.Fscanln(reader, &n)

	var str []byte
	fmt.Fscanf(reader, "%s", &str)

	for n != 1 {
		str[n-2] = table[check(str[n-2])][check(str[n-1])]
		n--
	}
	fmt.Printf("%c\n", str[0])
}
