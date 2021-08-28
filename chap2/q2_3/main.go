package main

import (
	"fmt"
	"popcount/popcount"
	"strconv"
)

func main() {
	fmt.Println(strconv.Itoa(popcount.PopCount(100)))
	fmt.Println(strconv.Itoa(popcount.RoopPopCount(100)))
}
