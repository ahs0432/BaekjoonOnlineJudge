package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var num, addnum int
	fmt.Fscanln(reader, &num, &addnum)
	if isPrime(num) && isPrime(num+(addnum*1000000)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
