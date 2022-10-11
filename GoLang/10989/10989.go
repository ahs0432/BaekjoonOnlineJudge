package main

import (
	"bufio"
	"fmt"
	"os"
)

func quickSort(table []int, below int, upper int) {
	if below < upper {
		part := divide(table, below, upper)
		quickSort(table, below, upper-1)
		quickSort(table, part+1, upper)
	}
}

func divide(table []int, below int, upper int) int {
	i := below
	center := table[upper]

	for j := below; j < upper; j++ {
		if table[j] <= center {
			swap(&table[i], &table[j])
			i += 1
		}
	}

	swap(&table[i], &table[upper])
	return i
}

func swap(var1 *int, var2 *int) {
	tmp := *var1
	*var1 = *var2
	*var2 = tmp
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var num int
	fmt.Fscanln(reader, &num)

	table := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	quickSort(table, 0, num-1)

	for _, i := range table {
		fmt.Println(i)
	}
}
