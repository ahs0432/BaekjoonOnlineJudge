package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var str string
	fmt.Fscanln(reader, &str)
	tmp, _ := reader.ReadString('\n')

	for _, c := range strings.Split(strings.TrimSpace(tmp), " ") {
		str = strings.ReplaceAll(str, c, string(c[0]+32))
	}
	fmt.Println(str)
}
