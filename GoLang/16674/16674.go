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

	tmp := make([]int, 10)
	for _, c := range s {
		tmp[c-48]++
	}

	if tmp[3]+tmp[4]+tmp[5]+tmp[6]+tmp[7]+tmp[9] > 0 {
		fmt.Println(0)
	} else if tmp[2] > 0 && tmp[0] > 0 && tmp[1] > 0 && tmp[8] > 0 {
		if tmp[2] == tmp[0] && tmp[0] == tmp[1] && tmp[1] == tmp[8] {
			fmt.Println(8)
		} else {
			fmt.Println(2)
		}
	} else {
		fmt.Println(1)
	}
}
