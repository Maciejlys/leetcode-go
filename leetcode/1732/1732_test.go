package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		gain []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				gain: []int{-5, 1, 5, 0, -7},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				gain: []int{-4, -3, -2, -1, 4, 3, 2},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := largestAltitude(tt.args.gain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
