package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var num int
	fmt.Fscanln(reader, &num)

	var number []float64
	top := 0.0

	for i := 0; i < num; i++ {
		a := 0.0
		fmt.Fscanf(reader, "%f ", &a)

		if top < a {
			top = a
		}

		if number == nil {
			number = []float64{a}
		} else {
			number = append(number, a)
		}
	}

	total := 0.0

	for _, numb := range number {
		total = total + float64(numb)/float64(top)*100.0
	}

	fmt.Fprint(writer, total/float64(len(number)))
}
