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

	n, _ := reader.ReadString('\n')
	n = strings.TrimSuffix(n, "\n")
	s := strings.Split(n, " ")

	sum := 0
	for _, i := range s {
		tmp, _ := strconv.Atoi(i)
		sum += tmp
	}

	fmt.Println(sum)
}
