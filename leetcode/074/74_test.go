package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		matrix [][]int
		target int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
				target: 3,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
				target: 13,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := searchMatrix(tt.args.matrix, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
