package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var s float64 = 3.1415
	fmt.Printf(comma(s))
}

func comma(s float64) string {
	var buf bytes.Buffer
	var sign_s string

	if s >= 0 {
		sign_s = "+"
	} else {
		sign_s = "-"
		s = -1 * s
	}

	s_converted := strconv.FormatFloat(s, 'f', -1, 64)
	buf.WriteString(sign_s)
	buf.WriteString(s_converted[:2])
	buf.WriteString(",")
	buf.WriteString(s_converted[2:])
	return buf.String()
}
