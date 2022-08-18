package p1302

import (
	"github.com/ricky1993/solid-funicular/structures/tree"
	"testing"
)

func Test_deepestLeavesSum(t *testing.T) {
	type args struct {
		root *tree.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test#1",
			args: args{
				root: tree.BuildTree(tree.NewPtr(1), tree.NewPtr(2), tree.NewPtr(3), tree.NewPtr(4),
					tree.NewPtr(5), nil, tree.NewPtr(6), tree.NewPtr(7), nil, nil, nil, nil, tree.NewPtr(8)),
			},
			want: 15,
		},
		{
			name: "test#2",
			args: args{
				root: tree.BuildTree(tree.NewPtr(6), tree.NewPtr(7), tree.NewPtr(8), tree.NewPtr(2),
					tree.NewPtr(7), tree.NewPtr(1), tree.NewPtr(3), tree.NewPtr(9), nil,
					tree.NewPtr(1), tree.NewPtr(4), nil, nil, nil, tree.NewPtr(5)),
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepestLeavesSum(tt.args.root); got != tt.want {
				t.Errorf("deepestLeavesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
