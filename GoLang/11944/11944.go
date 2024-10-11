package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	var str string = ""

	for i := 0; i < n; i++ {
		str += strconv.Itoa(n)
	}
	if len(str) < m {
		fmt.Println(str)
	} else {
		fmt.Println(str[:m])
	}
}
