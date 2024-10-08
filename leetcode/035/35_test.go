package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 3, 5, 6}, target: 5,
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1, 3, 5, 6}, target: 2,
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{1, 3, 5, 6}, target: 7,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := searchInsert(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
