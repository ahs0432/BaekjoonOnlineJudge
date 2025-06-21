package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	subjects := map[string]string{
		"Algorithm":              "204",
		"DataAnalysis":           "207",
		"ArtificialIntelligence": "302",
		"CyberSecurity":          "B101",
		"Network":                "303",
		"Startup":                "501",
		"TestStrategy":           "105",
	}

	var n int
	fmt.Fscanln(reader, &n)

	var tmp string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintln(writer, subjects[tmp])
	}
	writer.Flush()
}
