package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	sum := 0
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)
		d, _ := strconv.Atoi(strings.Split(tmp, "-")[1])

		if d <= 90 {
			sum++
		}
	}

	fmt.Println(sum)
}
