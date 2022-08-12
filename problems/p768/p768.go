package p768

import "sort"

func maxChunksToSorted(arr []int) int {
	newArr := make([]int, len(arr))
	sumArr := make([]int, len(arr))
	copy(newArr, arr)
	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i] < newArr[j]
	})
	for i := range newArr {
		if i == 0 {
			sumArr[i] = newArr[i]
		} else {
			sumArr[i] = newArr[i] + sumArr[i-1]
		}
	}
	ret := 0
	sum := 0
	for i := range arr {
		sum += arr[i]
		if sum == sumArr[i] {
			ret++
		}
	}
	return ret
}
