package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			},
			want: -1,
		},
		{
			name: "example 3",
			args: args{
				grid: [][]int{{0, 2}},
			},
			want: 0,
		},
		{
			name: "example 4",
			args: args{
				grid: [][]int{{0, 2, 2}},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := orangesRotting(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
