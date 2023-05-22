package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	if m-n > 0 {
		if m-n <= 20 {
			fmt.Println("You are speeding and your fine is $100.")
		} else if m-n <= 30 {
			fmt.Println("You are speeding and your fine is $270.")
		} else {
			fmt.Println("You are speeding and your fine is $500.")
		}
	} else {
		fmt.Println("Congratulations, you are within the speed limit!")
	}
}
