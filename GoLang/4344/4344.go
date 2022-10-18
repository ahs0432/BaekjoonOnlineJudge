package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	var num1 int
	fmt.Fscanln(reader, &num1)

	for i := 0; i < num1; i++ {
		var num2 int
		fmt.Fscanf(reader, "%d ", &num2)

		var number []int
		var total int
		for j := 0; j < num2; j++ {
			var num3 int
			fmt.Fscanf(reader, "%d ", &num3)

			if number == nil {
				number = []int{num3}
			} else {
				number = append(number, num3)
			}

			total += num3
		}

		avg := total / num2
		mincount := 0

		for _, j := range number {
			if avg < j {
				mincount += 1
			}
		}

		fmt.Fprintf(writer, "%.3f%s\n", math.Round(float64((float64(mincount)/float64(len(number)))*100000))/1000, "%")
	}
}
