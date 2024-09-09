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
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{2, 1, 1, 2},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := rob(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
