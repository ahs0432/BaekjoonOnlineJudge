package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	check := true

	var n1 int
	fmt.Fscanln(reader, &n1)

	if !(n1 == 9 || n1 == 8) {
		check = false
	}

	var n2 int
	fmt.Fscanln(reader, &n2)

	var n3 int
	fmt.Fscanln(reader, &n3)

	if check && n2 != n3 {
		check = false
	}

	var n4 int
	fmt.Fscanln(reader, &n4)
	if check && !(n4 == 9 || n4 == 8) {
		check = false
	}

	if check {
		fmt.Println("ignore")
	} else {
		fmt.Println("answer")
	}
}
