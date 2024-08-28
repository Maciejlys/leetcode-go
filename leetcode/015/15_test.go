package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "example 3",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
