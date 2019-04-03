package main

import (
	"fmt"
	"ica01/opg3/sum"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 {
		if i1, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil {
			if i2, err := strconv.ParseInt(os.Args[2], 10, 64); err == nil {
				sum := sum.SumInt64(i1, i2)
				fmt.Printf("%T, %v + %v = %v\n", i1, i1, i2, sum)
			}
		}
		if i1, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
			if i2, err := strconv.ParseFloat(os.Args[2], 64); err == nil {
				sum := sum.SumFloat64(i1, i2)
				fmt.Printf("%T, %v + %v = %v\n", i1, i1, i2, sum)
			}
		}
	} else {
		fmt.Println("Please provide 2 arguments of max size int 64 or float 64.")
	}
}
