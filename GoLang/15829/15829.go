package main

import (
	"bufio"
	"fmt"
	"os"
)

var mod int = 1234567891
var mul int = 31

func main() {
	reader := bufio.NewReader(os.Stdin)

	var num int
	var str string

	fmt.Fscanln(reader, &num)
	fmt.Fscanln(reader, &str)

	answer := 0
	r := 1

	for i := 0; i < num; i++ {
		answer = (answer + (int(str[i]-'a')+1)*r) % mod
		r = (r * mul) % mod
	}

	fmt.Println(answer)
}
