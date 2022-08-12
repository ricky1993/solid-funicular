package p769

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test#1",
			args: args{
				arr: []int{4, 3, 2, 1, 0},
			},
			want: 1,
		},
		{
			name: "test#2",
			args: args{
				arr: []int{1, 0, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "test#3",
			args: args{
				arr: []int{2, 0, 1},
			},
			want: 1,
		},
		{
			name: "test#4",
			args: args{
				arr: []int{0},
			},
			want: 1,
		},
		{
			name: "test#4",
			args: args{
				arr: []int{0, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunksToSorted(tt.args.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
