package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var v1, v2, field string
	fmt.Fscanln(reader, &v1, &field, &v2)

	if field == "AND" {
		if v1 == "true" && v2 == "true" {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	} else if field == "OR" {
		if v1 == "true" || v2 == "true" {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
