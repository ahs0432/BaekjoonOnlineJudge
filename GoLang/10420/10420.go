package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	ans := time.Date(2014, 4, 2, 0, 0, 0, 0, time.Local).AddDate(0, 0, n-1)
	fmt.Println(ans.Format("2006-01-02"))
}
