package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")
	c := strings.Split(s, " ")

	money := 5000
	for _, n := range c {
		if n == "1" {
			money -= 500
		} else if n == "2" {
			money -= 800
		} else if n == "3" {
			money -= 1000
		}
	}
	fmt.Println(money)
}
