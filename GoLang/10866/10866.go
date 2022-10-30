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

	queue := make([]int, 0)

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(reader, &word)
		if word == "push_front" {
			var num int
			fmt.Fscanln(reader, &num)
			queue = append([]int{num}, queue...)
		} else if word == "push_back" {
			var num int
			fmt.Fscanln(reader, &num)
			queue = append(queue, num)
		} else if word == "front" {
			if len(queue) != 0 {
				fmt.Println(queue[0])
			} else {
				fmt.Println(-1)
			}
		} else if word == "back" {
			if len(queue) != 0 {
				fmt.Println(queue[len(queue)-1])
			} else {
				fmt.Println(-1)
			}
		} else if word == "size" {
			fmt.Println(len(queue))
		} else if word == "empty" {
			if len(queue) == 0 {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		} else if word == "pop_front" {
			if len(queue) == 0 {
				fmt.Println(-1)
			} else {
				fmt.Println(queue[0])
				queue = queue[1:]
			}
		} else if word == "pop_back" {
			if len(queue) == 0 {
				fmt.Println(-1)
			} else {
				fmt.Println(queue[len(queue)-1])
				queue = queue[:len(queue)-1]
			}
		}
	}
}
