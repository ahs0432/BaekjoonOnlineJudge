package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0
	min := 2147483647
	check := 0

	for i := 0; i < 3; i++ {
		var tmp int
		if i == 2 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}
		sum += tmp

		if min > tmp {
			min = tmp
			check = i
		}
	}

	if sum < 100 {
		if check == 0 {
			fmt.Println("Soongsil")
		} else if check == 1 {
			fmt.Println("Korea")
		} else {
			fmt.Println("Hanyang")
		}
	} else {
		fmt.Println("OK")
	}
}
