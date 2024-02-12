package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b float64
	fmt.Fscanln(reader, &a, &b)

	ta := -a + math.Sqrt(math.Pow(a, 2)-b)
	tb := -a - math.Sqrt(math.Pow(a, 2)-b)

	if ta > tb {
		fmt.Println(int(tb), int(ta))
	} else if ta == tb {
		fmt.Println(int(ta))
	} else {
		fmt.Println(int(ta), int(tb))
	}
}
