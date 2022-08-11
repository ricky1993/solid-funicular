package array

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode[int]
	}{
		{
			name: "test#1",
			args: args{
				values: []int{1, 2, 3},
			},
			want: &ListNode[int]{
				Val: 1,
				Next: &ListNode[int]{
					Val: 2,
					Next: &ListNode[int]{
						Val: 3,
					},
				},
			},
		},
		{
			name: "test#2",
			args: args{
				values: []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}
