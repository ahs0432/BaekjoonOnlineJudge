package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n1, n2, n3 int
	fmt.Fscanln(reader, &n1)
	fmt.Fscanln(reader, &n2)
	fmt.Fscanln(reader, &n3)

	if n1+n2+n3 != 180 {
		fmt.Println("Error")
	} else if n1 == 60 && n2 == 60 && n3 == 60 {
		fmt.Println("Equilateral")
	} else if n1 == n2 || n2 == n3 || n1 == n3 {
		fmt.Println("Isosceles")
	} else {
		fmt.Println("Scalene")
	}
}
