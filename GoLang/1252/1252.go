package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b []byte
	fmt.Fscanln(reader, &a, &b)

	if len(a) < len(b) {
		a, b = b, a
	}

	for len(a) != len(b) {
		b = append([]byte{'0'}, b...)
	}

	sum := []byte{}
	aTmp, bTmp := 0, 0
	prev := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' {
			aTmp = 1
		} else {
			aTmp = 0
		}

		if b[i] == '1' {
			bTmp = 1
		} else {
			bTmp = 0
		}

		if aTmp+bTmp+prev == 0 {
			prev = 0
			sum = append([]byte{'0'}, sum...)
		} else if aTmp+bTmp+prev == 1 {
			prev = 0
			sum = append([]byte{'1'}, sum...)
		} else if aTmp+bTmp+prev == 2 {
			prev = 1
			sum = append([]byte{'0'}, sum...)
		} else if aTmp+bTmp+prev == 3 {
			prev = 1
			sum = append([]byte{'1'}, sum...)
		}
	}

	if prev == 1 {
		sum = append([]byte{'1'}, sum...)
	}

	start := -1

	for i := 0; i < len(sum); i++ {
		if sum[i] == '1' {
			start = i
			break
		}
	}

	if start == -1 {
		fmt.Fprintln(writer, 0)
	} else {
		for ; start < len(sum); start++ {
			fmt.Fprint(writer, string(sum[start]))
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
