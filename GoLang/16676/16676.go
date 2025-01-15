package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	c := 1

	if n <= 10 {
		fmt.Println(1)
	} else {
		var tmp string
		for n >= c {
			tmp = strconv.Itoa(c)
			tmp += "1"
			c, _ = strconv.Atoi(tmp)
		}
		fmt.Println(len(tmp[:len(tmp)-1]))
	}
}
