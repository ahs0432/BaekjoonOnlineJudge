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

	var s string
	fmt.Fscanln(reader, &s)

	ts := strings.Split(s, ":")
	ti := make([]int, 3)

	for i := 0; i < 3; i++ {
		ti[i], _ = strconv.Atoi(ts[i])
	}

	cnt := 0
	if (ti[0] > 0 && ti[0] < 13) && (ti[1] >= 0 && ti[1] < 60) && (ti[2] >= 0 && ti[2] < 60) {
		cnt += 2
	}

	if (ti[1] > 0 && ti[1] < 13) && (ti[0] >= 0 && ti[0] < 60) && (ti[2] >= 0 && ti[2] < 60) {
		cnt += 2
	}

	if (ti[2] > 0 && ti[2] < 13) && (ti[1] >= 0 && ti[1] < 60) && (ti[0] >= 0 && ti[0] < 60) {
		cnt += 2
	}

	fmt.Println(cnt)
}
