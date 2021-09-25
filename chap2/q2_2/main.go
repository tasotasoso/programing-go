package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"q2_2/tempconv"
)

func main() {

	var inputs []string

	if len(os.Args[1:]) == 0 {
		fmt.Print("Enter numbers with , separator")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		in := scanner.Text()
		inputs = strings.Split(in, ",")
		fmt.Println("in: ", inputs)
	} else {
		inputs = os.Args[1:]
	}

	for _, arg := range inputs {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		ab := tempconv.AbsoluteZeroC(t)

		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", ab, tempconv.KToC(ab), c, tempconv.CToK(c))
	}

}
