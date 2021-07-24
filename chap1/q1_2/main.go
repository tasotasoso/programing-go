package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		var s string = ""
		var sep string = " "
		s += strconv.Itoa(i) + sep + os.Args[i]
		fmt.Println(s)
	}
}
