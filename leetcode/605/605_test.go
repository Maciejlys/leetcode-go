package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		flowerbed []int
		n         int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         2,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
