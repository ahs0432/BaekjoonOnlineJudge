package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s1, s2 string
	fmt.Fscanln(reader, &s1)
	fmt.Fscanln(reader, &s2)

	var add, res string
	for i := 0; i < len(s1); i++ {
		add = add + string(s1[i]) + string(s2[i])
	}

	for len(add) != 2 {
		for i := 0; i < len(add)-1; i++ {
			res = res + string('0'+((add[i]-'0')+(add[i+1]-'0'))%10)
		}

		add = res
		res = ""
	}

	fmt.Println(add)
}
