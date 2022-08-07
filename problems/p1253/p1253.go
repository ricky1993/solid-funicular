// https://leetcode.com/problems/reconstruct-a-2-row-binary-matrix/
package p1253

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	ret := [][]int{}
	upperList := make([]int, len(colsum))
	lowerList := make([]int, len(colsum))
	for i := range colsum {
		if colsum[i] == 2 {
			upperList[i] = 1
			lowerList[i] = 1
			upper -= 1
			lower -= 1
		}
	}
	for i := range colsum {
		if colsum[i] == 1 {
			if upper > 0 {
				upperList[i] = 1
				upper -= 1
			} else if lower > 0 {
				lowerList[i] = 1
				lower -= 1
			} else {
				return ret
			}
		}
	}
	if lower != 0 || upper != 0 {
		return ret
	}
	ret = append(ret, upperList, lowerList)
	return ret
}
