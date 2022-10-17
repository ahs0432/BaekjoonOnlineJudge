package main

import (
	"bufio"
	"fmt"
	"os"
)

func pop(table []int) []int {
	return table[:len(table)-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscanln(reader, &num)

	stored := []int{}
	answer := []string{}
	cur := 1

	for i := 0; i < num; i++ {
		var num int
		fmt.Fscanln(reader, &num)

		for ; cur <= num; cur++ {
			answer = append(answer, "+")
			stored = append(stored, cur)
		}

		if num == stored[len(stored)-1] {
			answer = append(answer, "-")
			stored = pop(stored)
		} else {
			fmt.Println("NO")
			return
		}
	}

	for _, a := range answer {
		fmt.Println(a)
	}
}
