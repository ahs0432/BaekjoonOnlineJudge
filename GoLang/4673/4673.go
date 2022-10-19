package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	//	reader := bufio.NewReader(os.Stdin)
	defer writer.Flush()

	selfNumList := selfNumber(1, 10000, [10000]bool{})

	for i, selfNum := range selfNumList {
		if !selfNum {
			fmt.Fprintln(writer, i+1)
		}
	}

}

func selfNumber(num int, end int, selfNumList [10000]bool) [10000]bool {
	if num == end {
		return selfNumList
	}

	if (numUnit(num)+num-1 <= end-1) && !(selfNumList[numUnit(num)+num-1]) {
		selfNumList[numUnit(num)+num-1] = true
	}

	selfNumList = selfNumber(num+1, end, selfNumList)

	return selfNumList
}

func numUnit(num int) int {
	if num >= 10 {
		a := numUnit(int(num / 10))
		return a + num%10
	} else {
		return num
	}
}
