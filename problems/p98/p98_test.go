package p98

import (
	"github.com/ricky1993/solid-funicular/structures/tree"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *tree.TreeNode[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				root: &tree.TreeNode[int]{Val: 2},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				root: &tree.TreeNode[int]{
					Val:  2,
					Left: &tree.TreeNode[int]{Val: 1},
				},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				root: &tree.TreeNode[int]{
					Val:   2,
					Left:  &tree.TreeNode[int]{Val: 1},
					Right: &tree.TreeNode[int]{Val: 3},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
