package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	month := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var m, d int
	fmt.Fscanln(reader, &m, &d)

	day := 0
	for i := 0; i < m-1; i++ {
		day += month[i]
	}
	day = (day + d) % 7

	if day == 0 {
		fmt.Println("SUN")
	} else if day == 1 {
		fmt.Println("MON")
	} else if day == 2 {
		fmt.Println("TUE")
	} else if day == 3 {
		fmt.Println("WED")
	} else if day == 4 {
		fmt.Println("THU")
	} else if day == 5 {
		fmt.Println("FRI")
	} else if day == 6 {
		fmt.Println("SAT")
	}
}
