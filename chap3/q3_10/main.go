package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s string = "abcde"
	fmt.Printf(comma(s))
}

func comma(s string) string {
	var buf bytes.Buffer
	buf.WriteString(s[:2])
	buf.WriteString(",")
	buf.WriteString(s[2:])
	return buf.String()
}
