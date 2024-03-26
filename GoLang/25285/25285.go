package main

import (
	"bufio"
	"fmt"
	"os"
)

func bmi(cm, kg float64) float64 {
	return kg / ((cm / 100) * (cm / 100))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var cm, kg float64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &cm, &kg)

		if cm < 140.1 {
			fmt.Fprintln(writer, 6)
		} else if cm < 146 {
			fmt.Fprintln(writer, 5)
		} else if cm < 159 {
			fmt.Fprintln(writer, 4)
		} else if cm < 161 {
			bmi := bmi(cm, kg)

			if bmi < 16 || bmi >= 35 {
				fmt.Fprintln(writer, 4)
			} else {
				fmt.Fprintln(writer, 3)
			}
		} else if cm < 204 {
			bmi := bmi(cm, kg)

			if bmi < 16 || bmi >= 35 {
				fmt.Fprintln(writer, 4)
			} else if (bmi >= 16 && bmi < 18.5) || (bmi >= 30 && bmi < 35) {
				fmt.Fprintln(writer, 3)
			} else if (bmi >= 18.5 && bmi < 20) || (bmi >= 25 && bmi < 30) {
				fmt.Fprintln(writer, 2)
			} else {
				fmt.Fprintln(writer, 1)
			}
		} else {
			fmt.Fprintln(writer, 4)
		}
	}

	writer.Flush()
}
