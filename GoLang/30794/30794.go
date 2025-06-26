package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var lv string
	fmt.Fscanln(reader, &n, &lv)

	if lv == "miss" {
		fmt.Println(n * 0)
	} else if lv == "bad" {
		fmt.Println(n * 200)
	} else if lv == "cool" {
		fmt.Println(n * 400)
	} else if lv == "great" {
		fmt.Println(n * 600)
	} else if lv == "perfect" {
		fmt.Println(n * 1000)
	}
}
