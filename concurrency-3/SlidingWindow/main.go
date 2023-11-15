package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{9, 4, 1, 7}
	k := 2
	result := minimumDifference(nums, k)
	fmt.Println("Min diff is", result)
}

func minimumDifference(nums []int, k int) int {
	//Introduced concurrency for fun :D
	//Sorts in parallel
	sortCh := make(chan struct{})
	go func() {
		sort.Ints(nums)
		sortCh <- struct{}{}
	}()
	minDiff := math.MaxInt32
	<-sortCh

	for i := 0; i+k-1 < len(nums); i++ {
		currentDiff := nums[i+k-1] - nums[i]
		minDiff = min(minDiff, currentDiff)
	}
	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
