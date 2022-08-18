package p1656

type OrderedStream struct {
	index int
	n     int
	num   []int
	s     []string
}

func Constructor(n int) OrderedStream {
	stream := OrderedStream{
		index: 0,
		n:     n,
		num:   make([]int, n),
		s:     make([]string, n),
	}
	return stream
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	idKey--
	this.num[idKey] = 1
	this.s[idKey] = value
	ret := make([]string, 0)
	for i := this.index; i < this.n; i++ {
		if this.num[i] == 0 {
			this.index = i
			break
		} else {
			ret = append(ret, this.s[i])
		}
	}
	return ret
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
