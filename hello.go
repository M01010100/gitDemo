package main

import (
	"fmt"
	"strings"
)

var output string

func printOut(in_1 string, in_2 string, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(in_1)
		sb.WriteString(" ")
		sb.WriteString(in_2)
		sb.WriteString("\n")
	}
	return sb.String()
}
func main() {
	output := printOut("Pesto is", "a Penguin", 5)
	fmt.Println(output)
}
