package p769

func maxChunksToSorted(arr []int) int {
	ret := 0
	sum := 0
	for i := range arr {
		sum += arr[i]
		if sum == i*(i+1)/2 {
			ret++
		}
	}
	return ret
}
