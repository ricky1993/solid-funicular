package tree

// TreeNode is defined for a binary tree node.
type TreeNode[T comparable] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func BuildTree[T comparable](values ...*T) *TreeNode[T] {
	if len(values) == 0 {
		return nil
	}
	rootVal := values[0]
	if rootVal == nil {
		panic("root node cannot be nil")
	}
	rootNode := &TreeNode[T]{
		Val:   *rootVal,
		Left:  nil,
		Right: nil,
	}
	var rootChan = make(chan *TreeNode[T], (len(values)+1)/2)
	i := 1
	rootChan <- rootNode
	for current := range rootChan {
		if i == len(values) {
			break
		}
		val := values[i]
		if val != nil {
			current.Left = &TreeNode[T]{
				Val: *val,
			}
			rootChan <- current.Left
		}
		i++
		if i == len(values) {
			break
		}
		val = values[i]
		if val != nil {
			current.Right = &TreeNode[T]{
				Val: *val,
			}
			rootChan <- current.Right
		}
		i++
	}
	close(rootChan)
	return rootNode

}

func NewPtr[T comparable](in T) *T {
	out := in
	return &out
}
