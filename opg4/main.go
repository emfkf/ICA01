package main

import (
	"fmt"
	"ica01/opg4/algorithms"
	"math/rand"
)

func main() {
	arr := perm(5)
	fmt.Println(arr)
	algorithms.BubbleSortMod(arr)
	fmt.Println(arr)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}
