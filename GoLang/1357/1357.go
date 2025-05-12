package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y string
	fmt.Fscanln(reader, &x, &y)

	length := len(x)
	if len(x) < len(y) {
		length = len(y)
	}

	ans := 0
	now := 1
	for i := 0; i < length; i++ {
		if len(x)-1 < i {
			ans += int(y[i]-'0') * now
		} else if len(y)-1 < i {
			ans += int(x[i]-'0') * now
		} else {
			ans += (int(x[i]-'0') + int(y[i]-'0')) * now
		}
		now *= 10
	}

	res := ""
	for ans != 0 {
		res += strconv.Itoa(ans % 10)
		ans /= 10
	}
	ans, _ = strconv.Atoi(res)
	fmt.Fprintln(writer, ans)
	writer.Flush()
}
