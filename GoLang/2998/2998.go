package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	for len(s)%3 != 0 {
		s = "0" + s
	}

	ans := ""
	now := ""
	for i := 0; i < len(s); i += 3 {
		now = s[i : i+3]
		if now == "000" {
			ans += "0"
		} else if now == "001" {
			ans += "1"
		} else if now == "010" {
			ans += "2"
		} else if now == "011" {
			ans += "3"
		} else if now == "100" {
			ans += "4"
		} else if now == "101" {
			ans += "5"
		} else if now == "110" {
			ans += "6"
		} else if now == "111" {
			ans += "7"
		}
	}
	fmt.Println(ans)
}
