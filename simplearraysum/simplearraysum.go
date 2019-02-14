package simplearraysum

// SimpleArraySum - Given an array of integers, find the sum of its elements.
func SimpleArraySum(ar []int32) int32 {
	var sum int32

	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}

	return sum
}
