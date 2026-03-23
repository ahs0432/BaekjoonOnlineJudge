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
	var d string
	fmt.Fscanln(reader, &d, &n)
	day, _ := time.Parse("2006-01-02", d)
	result := day.AddDate(0, 0, n-1)
	fmt.Println(result.Format("2006-01-02"))
}
