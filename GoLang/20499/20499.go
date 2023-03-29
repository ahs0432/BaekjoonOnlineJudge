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

	var n string
	fmt.Fscanln(reader, &n)

	table := strings.Split(n, "/")
	tInt := make([]int, 3)

	tInt[0], _ = strconv.Atoi(table[0])
	tInt[1], _ = strconv.Atoi(table[1])
	tInt[2], _ = strconv.Atoi(table[2])

	if tInt[1] == 0 || tInt[0]+tInt[2] < tInt[1] {
		fmt.Println("hasu")
	} else {
		fmt.Println("gosu")
	}
}
