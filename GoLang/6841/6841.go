package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		var n string
		fmt.Fscanln(reader, &n)

		if n == "TTYL" {
			fmt.Fprintln(writer, "talk to you later")
			break
		} else if n == "CU" {
			fmt.Fprintln(writer, "see you")
		} else if n == ":-)" {
			fmt.Fprintln(writer, "I’m happy")
		} else if n == ":-(" {
			fmt.Fprintln(writer, "I’m unhappy")
		} else if n == ";-)" {
			fmt.Fprintln(writer, "wink")
		} else if n == ":-P" {
			fmt.Fprintln(writer, "stick out my tongue")
		} else if n == "(~.~)" {
			fmt.Fprintln(writer, "sleepy")
		} else if n == "TA" {
			fmt.Fprintln(writer, "totally awesome")
		} else if n == "CCC" {
			fmt.Fprintln(writer, "Canadian Computing Competition")
		} else if n == "CUZ" {
			fmt.Fprintln(writer, "because")
		} else if n == "TY" {
			fmt.Fprintln(writer, "thank-you")
		} else if n == "YW" {
			fmt.Fprintln(writer, "you’re welcome")
		} else {
			fmt.Fprintln(writer, n)
		}
	}

	writer.Flush()
}
