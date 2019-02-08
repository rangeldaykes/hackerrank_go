package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	fmt.Println("inicio")

	arr := []int32{1, 3, 5, 7, 9}

	//miniMaxSum(arr)
	solution2(arr)

}

func solution1(arr []int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	var ValueMin, ValueMax int64
	for i := 0; i < 4; i++ {
		ValueMin += int64(arr[i])
		ValueMax += int64(arr[i+1])
	}

	fmt.Printf("%v %v", ValueMin, ValueMax)
}

func solution2(arr []int32) {

	var ArrTotal, ArrMin, ArrMax int64
	ArrTotal, ArrMin, ArrMax = sum(arr)

	fmt.Printf("%v", ArrTotal)
	fmt.Println("")
	fmt.Printf("%v %v", ArrTotal-ArrMax, ArrTotal-ArrMin)
}

func sum(arr []int32) (sum, min, max int64) {
	min = math.MaxInt64

	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
		if int64(arr[i]) < min {
			min = int64(arr[i])
		}
		if int64(arr[i]) > max {
			max = int64(arr[i])
		}
	}

	return
}
