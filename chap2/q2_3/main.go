package main

import (
	"fmt"
	"popcount/popcount"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strconv.Itoa(popcount.PopCount(100)))
	end := time.Now()
	fmt.Printf("最初のバージョンは%f秒\n", (end.Sub(start)).Seconds())

	start = time.Now()
	fmt.Println(strconv.Itoa(popcount.RoopPopCount(100)))
	end = time.Now()
	fmt.Printf("修正後のバージョンは%f秒\n", (end.Sub(start)).Seconds())
}
