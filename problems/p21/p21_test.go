package p21

import (
	"github.com/ricky1993/solid-funicular/structures/array"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *array.ListNode[int]
		list2 *array.ListNode[int]
	}
	tests := []struct {
		name string
		args args
		want *array.ListNode[int]
	}{
		{
			name: "test#1",
			args: args{
				list1: array.NewList(1, 2, 4),
				list2: array.NewList(1, 3, 4),
			},
			want: array.NewList(1, 1, 2, 3, 4, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
