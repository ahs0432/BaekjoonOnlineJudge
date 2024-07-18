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

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)
		fmt.Println(string(tmp[0]) + string(tmp[len(tmp)-1]))
	}
}
