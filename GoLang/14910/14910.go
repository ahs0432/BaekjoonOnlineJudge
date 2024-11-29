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

	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	c := strings.Split(s, " ")

	prev := -2147483648
	tmp := 0
	for _, i := range c {
		tmp, _ = strconv.Atoi(i)
		if prev > tmp {
			fmt.Println("Bad")
			return
		}
		prev = tmp
	}
	fmt.Println("Good")
}
