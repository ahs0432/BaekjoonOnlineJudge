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

	count := 0

	if num < 100 {
		count = num
	} else if num < 111 {
		count = 99
	} else {
		count = 99
		for i := 111; i <= num; i++ {
			numList := numUnit(i)

			diff := 0
			first := true
			diffCheck := true
			for j := 0; j < len(numList)-1; j++ {
				if first {
					first = false
					diff = numList[j] - numList[j+1]
				} else if diff != numList[j]-numList[j+1] {
					diffCheck = false
					break
				}
			}

			if diffCheck {
				count++
			}
		}
	}
	fmt.Fprintln(writer, count)
}

func numUnit(num int) []int {
	if num >= 10 {
		a := numUnit(int(num / 10))
		a = append(a, num%10)
		return a
	} else {
		return []int{num}
	}
}
