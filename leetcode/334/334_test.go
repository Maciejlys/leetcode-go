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
		want bool
	}{
		{
			name: "example 1",
			args: args{
				[]int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				[]int{5, 4, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				[]int{2, 1, 5, 0, 4, 6},
			},
			want: true,
		},
		{
			name: "example 4",
			args: args{
				[]int{20, 100, 10, 12, 5, 13},
			},
			want: true,
		},
		{
			name: "example 5",
			args: args{
				[]int{2, 4, -2, -3},
			},
			want: false,
		},
		{
			name: "example 6",
			args: args{
				[]int{1, 5, 0, 4, 1, 3},
			},
			want: true,
		},
		{
			name: "example 7",
			args: args{
				[]int{0, 4, 2, 1, 0, -1, -3},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := increasingTriplet(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
