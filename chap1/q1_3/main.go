package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < 100000; i++ {
		s += sep + strconv.Itoa(i)
		sep = " "
	}
	end := time.Now()

	fmt.Printf("非効率バージョンは%f秒\n", (end.Sub(start)).Seconds())

	var str_arr [100000]string
	for i := 1; i < 100000; i++ {
		str_arr[i] = strconv.Itoa(i)
	}

	start = time.Now()
	s = strings.Join(str_arr[:], " ")
	end = time.Now()

	fmt.Printf("効率的なバージョンは%f秒\n", (end.Sub(start)).Seconds())
}
