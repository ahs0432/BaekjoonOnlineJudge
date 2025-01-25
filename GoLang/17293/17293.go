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

	for i := n; i >= 0; i-- {
		if i > 2 {
			fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", i, i)
			fmt.Printf("Take one down and pass it around, %d bottles of beer on the wall.\n\n", i-1)
		} else if i == 2 {
			fmt.Printf("2 bottles of beer on the wall, 2 bottles of beer.\n")
			fmt.Printf("Take one down and pass it around, 1 bottle of beer on the wall.\n\n")
		} else if i == 1 {
			fmt.Printf("1 bottle of beer on the wall, 1 bottle of beer.\n")
			fmt.Printf("Take one down and pass it around, no more bottles of beer on the wall.\n\n")
		}
	}

	fmt.Printf("No more bottles of beer on the wall, no more bottles of beer.\n")
	if n == 1 {
		fmt.Printf("Go to the store and buy some more, 1 bottle of beer on the wall.\n")
	} else {
		fmt.Printf("Go to the store and buy some more, %d bottles of beer on the wall.\n", n)
	}
}
