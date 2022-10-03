package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	var num1 int
	fmt.Fscanln(reader, &num1)
	defer writer.Flush()

	if num1 == 0 {
		fmt.Printf("1")
	} else if num1 < 10 {
		data := (num1 * 10) + num1
		count := 1

		for num1 != data {
			data = calc(data)
			count++
		}

		fmt.Fprint(writer, strconv.Itoa(count))
	} else {
		data := calc(num1)
		count := 1

		for num1 != data {
			data = calc(data)
			count++
		}

		fmt.Fprint(writer, strconv.Itoa(count))
	}
}

func calc(a int) int {
	return ((a % 10) * 10) + ((int(a/10) + (a % 10)) % 10)
}
