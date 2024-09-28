package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	joi := 0
	ioi := 0
	for i := 0; i < len(s)-2; i++ {
		if s[i:i+3] == "JOI" {
			joi++
		} else if s[i:i+3] == "IOI" {
			ioi++
		}
	}
	fmt.Println(joi)
	fmt.Println(ioi)
}
