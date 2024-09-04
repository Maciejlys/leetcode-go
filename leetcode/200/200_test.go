package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		grid [][]byte
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := numIslands(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
