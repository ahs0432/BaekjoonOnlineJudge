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
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var line string
	for i := 0; i < t; i++ {
		line, _ = reader.ReadString('\n')
		x := strings.Fields(line)
		n, _ := strconv.Atoi(x[0])

		k := make([][]int, n)
		for j := 0; j < n; j++ {
			line, _ = reader.ReadString('\n')
			nums := strings.Fields(line)
			row := make([]int, len(nums))
			for idx, val := range nums {
				row[idx], _ = strconv.Atoi(val)
			}
			k[j] = row
		}

		line, _ = reader.ReadString('\n')
		stickerStrings := strings.Fields(line)
		stickerList := make([]int, len(stickerStrings))
		for idx, val := range stickerStrings {
			stickerList[idx], _ = strconv.Atoi(val)
		}

		ans := 0
		for j := 0; j < n; j++ {
			m := 100
			numIndices := k[j][0]
			for l := 1; l <= numIndices; l++ {
				stickerIdx := k[j][l] - 1
				if stickerList[stickerIdx] < m {
					m = stickerList[stickerIdx]
				}
			}
			ans += m * k[j][len(k[j])-1]
		}

		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
