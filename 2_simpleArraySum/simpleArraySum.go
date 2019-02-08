package main

import (
	"fmt"
)

func main() {

	fmt.Println("inicio")

	arr := []int32{1, 3, 5, 7, 9}

	sum := simpleArraySum(arr)
	fmt.Println(sum)
}

func simpleArraySum(ar []int32) int32 {
	var sum int32

	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}

	return sum
}
