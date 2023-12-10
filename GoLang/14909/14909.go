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
	for _, c := range s {
		i, _ := strconv.Atoi(c)
		if i > 0 {
			sum++
		}
	}
	fmt.Println(sum)
}
