package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		candies []int
		extra   int
	}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "example 1",
			args: args{
				candies: []int{2, 3, 5, 1, 3},
				extra:   3,
			},
			want: []bool{true, true, true, false, true},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := kidsWithCandies(tt.args.candies, tt.args.extra); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
