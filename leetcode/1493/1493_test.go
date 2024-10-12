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
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 1, 0, 1},
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			},
			want: 5,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := longestSubarray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
