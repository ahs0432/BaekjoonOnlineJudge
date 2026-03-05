package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDivisorsSum(x int) int {
	sum := 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			sum += i
		}
	}
	return sum
}

func getDivisorsList(x int) []int {
	var table []int
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			table = append(table, i)
		}
	}
	return table
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		divisors := getDivisorsList(n)
		isGoodBye := true

		selfSum := getDivisorsSum(n)
		if selfSum-n <= n {
			fmt.Fprintln(writer, "BOJ 2022")
			isGoodBye = false
		} else {
			for j := 0; j < len(divisors)-1; j++ {
				d := divisors[j]
				if getDivisorsSum(d)-d > d {
					fmt.Fprintln(writer, "BOJ 2022")
					isGoodBye = false
					break
				}
			}
		}

		if isGoodBye {
			fmt.Fprintln(writer, "Good Bye")
		}
	}
	writer.Flush()
}
