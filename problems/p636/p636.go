// https://leetcode.com/problems/exclusive-time-of-functions/
package p636

import (
	"strconv"
	"strings"

	"github.com/ricky1993/solid-funicular/structures/array"
)

func exclusiveTime(n int, logs []string) []int {
	stack := array.Stack[[]string]{}
	ret := make([]int, n)
	for _, log := range logs {
		items := strings.Split(log, ":")
		if items[1] == "start" {
			if stack.Len() > 0 {
				top := *stack.Top()
				funcID, _ := strconv.Atoi(top[0])
				startTime, _ := strconv.Atoi(top[2])
				endTime, _ := strconv.Atoi(items[2])
				ret[funcID] += endTime - startTime
			}
			stack = append(stack, items)
		} else {
			// pop stack
			top := *stack.Top()
			stack.Pop()

			funcID, _ := strconv.Atoi(top[0])
			startTime, _ := strconv.Atoi(top[2])
			endTime, _ := strconv.Atoi(items[2])
			ret[funcID] += endTime - startTime + 1
			if stack.Len() > 0 {
				stack[stack.Len()-1][2] = strconv.Itoa(endTime + 1)
			}
		}
	}
	return ret
}
