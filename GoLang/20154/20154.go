package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := []int{3, 2, 1, 2, 3, 3, 3, 3, 1, 1, 3, 1, 3, 3, 1, 2, 2, 2, 1, 2, 1, 1, 2, 2, 2, 1}

	var k string
	fmt.Fscanln(reader, &k)

	ans := 0
	for i := 0; i < len(k); i++ {
		ans += table[k[i]-'A']
		ans %= 10
	}

	if ans%2 == 0 {
		fmt.Println("You're the winner?")
	} else {
		fmt.Println("I'm a winner!")
	}
}
