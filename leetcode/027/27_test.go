package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		val  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{3, 2, 2, 3}, val: 3,
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := removeElement(tt.args.nums, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
