package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var w float64
	fmt.Fscanln(reader, &w)
	fmt.Println(int(math.Sqrt(w*2) * 4))
}
