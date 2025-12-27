package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	grades := map[string]float64{
		"A+": 4.5,
		"A0": 4.0,
		"B+": 3.5,
		"B0": 3.0,
		"C+": 2.5,
		"C0": 2.0,
		"D+": 1.5,
		"D0": 1.0,
		"F":  0.0,
	}

	var totalCredits float64
	var totalScore float64

	var subject string
	var credits float64
	var grade string
	for i := 0; i < 20; i++ {
		fmt.Fscanln(reader, &subject, &credits, &grade)
		if grade == "P" {
			continue
		} else {
			totalScore += credits * grades[grade]
			totalCredits += credits
		}
	}
	fmt.Printf("%.6f\n", totalScore/totalCredits)
}
