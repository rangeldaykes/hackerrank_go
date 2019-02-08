package main

import (
	"fmt"
)

func main() {

	fmt.Println("inicio")

	arr := []int64{1, 3, 5, 7, 9}

	sum := aVeryBigSum(arr)
	fmt.Println(sum)
}

func aVeryBigSum(ar []int64) int64 {

	var sum int64
	for _, num := range ar {
		sum += num
	}
	return sum
}
