package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscanln(reader, &a, &b)
	fmt.Println(100-a, 100-b, (100 - (100 - a) - (100 - b)), ((100 - a) * (100 - b)), ((100-a)*(100-b))/100, ((100-a)*(100-b))%100)
	fmt.Println((100-(100-a)-(100-b))+(((100-a)*(100-b))/100), ((100-a)*(100-b))%100)
}
