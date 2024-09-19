package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		asteroids []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				asteroids: []int{5, 10, -5},
			},
			want: []int{5, 10},
		},
		{
			name: "example 2",
			args: args{
				asteroids: []int{8, -8},
			},
			want: []int{},
		},
		{
			name: "example 3",
			args: args{
				asteroids: []int{10, 2, -5},
			},
			want: []int{10},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
