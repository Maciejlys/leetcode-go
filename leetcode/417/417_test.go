package leetcode

import (
	"reflect"
	"testing"

	"github.com/Maciejlys/leetcode-go/util"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		heights [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				heights: [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			},
			want: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			name: "example 1",
			args: args{
				heights: [][]int{{10, 10, 10}, {10, 1, 10}, {10, 10, 10}},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := pacificAtlantic(tt.args.heights)
			want := tt.want
			util.Sort2DInts(got)
			util.Sort2DInts(want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
