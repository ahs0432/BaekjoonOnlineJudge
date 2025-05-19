package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	tableList := list.New()
	var table, res []int

	fmt.Fprintln(writer, "YES")
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			tableList.PushBack(i)
		} else {
			tableList.PushFront(i)
		}
	}

	i := 0
	index := 1
	for e := tableList.Front(); e != nil; e = e.Next() {
		if e.Value == 1 {
			index = i
		}

		table = append(table, e.Value.(int))
		fmt.Fprint(writer, e.Value, " ")
		i++
	}
	fmt.Fprintln(writer)

	for i := 1; i <= n; i++ {
		res = append(res, index+1)
		if i%2 == 0 {
			index = index - table[index]
		} else {
			index += table[index]
		}
	}

	for i := 0; i < len(res); i++ {
		fmt.Fprint(writer, res[i], " ")
	}
	fmt.Fprintln(writer)

	writer.Flush()
}
