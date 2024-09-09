package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		costs []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				costs: []int{10, 15, 20},
			},
			want: 15,
		},
		{
			name: "example 2",
			args: args{
				costs: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := minCostClimbingStairs(tt.args.costs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
