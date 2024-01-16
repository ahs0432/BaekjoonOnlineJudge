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

	for i := 0; i < num; i++ {
		var result string
		fmt.Fscanln(reader, &result)

		//resultSlice := strings.Split(result, "")
		var resultList []int
		var total int
		for j, r := range result {
			if resultList == nil {
				if string(r) == "O" {
					resultList = []int{1}
					total += 1
				} else {
					resultList = []int{0}
				}
			} else {
				if string(r) == "O" {
					resultList = append(resultList, (resultList[j-1] + 1))
					total += resultList[j]
				} else {
					resultList = append(resultList, 0)
				}

			}
		}

		fmt.Fprintln(writer, total)
	}

}
