package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var tmpNum, topNum int
	var tmpName, topName string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmpName, &tmpNum)

		if tmpNum > topNum {
			topNum = tmpNum
			topName = tmpName
		} else if tmpNum == topNum {
			if tmpName < topName {
				topNum = tmpNum
				topName = tmpName
			}
		}
	}
	fmt.Println(topName)
}
