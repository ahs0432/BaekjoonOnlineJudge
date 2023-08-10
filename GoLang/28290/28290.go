package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscanln(reader, &n)

	if n == "fdsajkl;" || n == "jkl;fdsa" {
		fmt.Println("in-out")
	} else if n == "asdf;lkj" || n == ";lkjasdf" {
		fmt.Println("out-in")
	} else if n == "asdfjkl;" {
		fmt.Println("stairs")
	} else if n == ";lkjfdsa" {
		fmt.Println("reverse")
	} else {
		fmt.Println("molu")
	}
}
