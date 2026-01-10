package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s []byte
	fmt.Fscanln(reader, &n, &s)

	if n == 1 {
		fmt.Println(string(s))
		for i := 0; i < len(s); i++ {
			if s[i] >= 'a' && s[i] <= 'z' {
				fmt.Print(string(s[i]))
			} else {
				fmt.Print("_" + string(s[i]+32))
			}
		}
		fmt.Println()
		fmt.Println(string(s[0]-32) + string(s[1:]))
	} else if n == 2 {
		for i := 0; i < len(s); i++ {
			if s[i] == '_' {
				fmt.Print(string(s[i+1] - 32))
				i++
			} else {
				fmt.Print(string(s[i]))
			}
		}
		fmt.Println()
		fmt.Println(string(s))
		fmt.Print(string(s[0] - 32))
		for i := 1; i < len(s); i++ {
			if s[i] == '_' {
				fmt.Print(string(s[i+1] - 32))
				i++
			} else {
				fmt.Print(string(s[i]))
			}
		}
		fmt.Println()
	} else if n == 3 {
		fmt.Println(string(s[0]+32) + string(s[1:]))
		fmt.Print(string(s[0] + 32))
		for i := 1; i < len(s); i++ {
			if s[i] >= 'a' && s[i] <= 'z' {
				fmt.Print(string(s[i]))
			} else {
				fmt.Print("_" + string(s[i]+32))
			}
		}
		fmt.Println()
		fmt.Println(string(s))
	}
}
