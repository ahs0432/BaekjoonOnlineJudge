package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var d string
	var n int
	fmt.Fscanln(reader, &d)
	fmt.Fscanln(reader, &n)

	ds := strings.Split(d, "-")
	yy, _ := strconv.Atoi(ds[0])
	mm, _ := strconv.Atoi(ds[1])
	dd, _ := strconv.Atoi(ds[2])

	dd += n
	mm += (dd - 1) / 30
	dd = (dd-1)%30 + 1
	yy += (mm - 1) / 12
	mm = (mm-1)%12 + 1

	fmt.Printf("%d-%02d-%02d\n", yy, mm, dd)
}
