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

	var s string
	fmt.Fscanln(reader, &s)

	var flag, cnt int

	for _, i := range s {
		if i == 'W' {
			cnt++
		} else {
			if cnt == 1 {
				flag = 2
			} else if cnt == 0 && flag == 0 {
				flag = 1
			} else if cnt == 0 && flag == 1 {
				flag = 0
			}
		}
	}

	if cnt >= 2 {
		if flag == 0 {
			fmt.Println(5)
		} else if flag == 1 {
			fmt.Println(1)
		} else {
			fmt.Println(6)
		}
	} else {
		fmt.Println(0)
	}
}
