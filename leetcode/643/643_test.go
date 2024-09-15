package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75000,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{5},
				k:    1,
			},
			want: 5,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{-1},
				k:    1,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := findMaxAverage(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
