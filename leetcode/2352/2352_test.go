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
				grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
			},
			want: 3,
		},
		{
			name: "example 3",
			args: args{
				grid: [][]int{{13, 13}, {13, 13}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := equalPairs(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
