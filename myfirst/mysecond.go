package main

import (
	"fmt"
)

func main() {
	x := 2
	y := 3
	sum := x + y
	arr := []int{3, 2, 1, 4}
	arr = append(arr, sum)
	fmt.Println(arr)
}

