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

	if n == "A+" {
		fmt.Println("4.3")
	} else if n == "A0" {
		fmt.Println("4.0")
	} else if n == "A-" {
		fmt.Println("3.7")
	} else if n == "B+" {
		fmt.Println("3.3")
	} else if n == "B0" {
		fmt.Println("3.0")
	} else if n == "B-" {
		fmt.Println("2.7")
	} else if n == "C+" {
		fmt.Println("2.3")
	} else if n == "C0" {
		fmt.Println("2.0")
	} else if n == "C-" {
		fmt.Println("1.7")
	} else if n == "D+" {
		fmt.Println("1.3")
	} else if n == "D0" {
		fmt.Println("1.0")
	} else if n == "D-" {
		fmt.Println("0.7")
	} else {
		fmt.Println("0.0")
	}
}
