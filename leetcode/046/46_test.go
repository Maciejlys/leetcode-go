package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func sum(nums []int) (sum int) {
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

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
				nums: []int{1, 2},
			},
			want: [][]int{{1, 2}, {2, 1}},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := permute(tt.args.nums)
			want := tt.want

			for i := range got {
				sort.Ints(got[i])
			}

			sort.Slice(got, func(i, j int) bool {
				return sum(got[i]) < sum(got[j])
			})

			for i := range tt.want {
				sort.Ints(want[i])
			}

			sort.Slice(want, func(i, j int) bool {
				return sum(want[i]) < sum(want[j])
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
