package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var A, a, B, b, P int
	fmt.Fscanln(reader, &A, &a, &B, &b, &P)

	if A > P || B > P {
		fmt.Println("No")
	} else if A+B <= P {
		fmt.Println("Yes")
	} else if a >= B {
		fmt.Println("Yes")
	} else if b >= A {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
