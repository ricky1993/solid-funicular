package p1656

import (
	"reflect"
	"testing"
)

func TestOrderedStream_Insert(t *testing.T) {
	type fields struct {
		n int
	}
	type args struct {
		idKey int
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
		want   [][]string
	}{
		{
			name: "test#1",
			fields: fields{
				n: 5,
			},
			args: []args{
				{
					idKey: 3,
					value: "ccccc",
				},
				{
					idKey: 1,
					value: "aaaaa",
				},
				{
					idKey: 2,
					value: "bbbbb",
				},
				{
					idKey: 5,
					value: "eeeee",
				},
				{
					idKey: 4,
					value: "ddddd",
				},
			},
			want: [][]string{{}, {"aaaaa"}, {"bbbbb", "ccccc"}, {}, {"ddddd", "eeeee"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.n)
			for i := range tt.args {
				if got := this.Insert(tt.args[i].idKey, tt.args[i].value); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("Insert() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
