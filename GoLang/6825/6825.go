package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w, h float64
	fmt.Fscanln(reader, &w)
	fmt.Fscanln(reader, &h)

	bmi := w / (h * h)

	if bmi >= 25 {
		fmt.Println("Overweight")
	} else if bmi >= 18 {
		fmt.Println("Normal weight")
	} else {
		fmt.Println("Underweight")
	}
}
