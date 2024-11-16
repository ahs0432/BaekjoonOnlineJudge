package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var d, m int
	fmt.Fscanln(reader, &d, &m)

	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	tDay := d
	for i := 0; i < m-1; i++ {
		tDay += months[i]
	}

	tDay = tDay % 7
	if tDay == 0 {
		fmt.Println("Wednesday")
	} else if tDay == 1 {
		fmt.Println("Thursday")
	} else if tDay == 2 {
		fmt.Println("Friday")
	} else if tDay == 3 {
		fmt.Println("Saturday")
	} else if tDay == 4 {
		fmt.Println("Sunday")
	} else if tDay == 5 {
		fmt.Println("Monday")
	} else if tDay == 6 {
		fmt.Println("Tuesday")
	}
}
