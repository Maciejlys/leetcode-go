package leetcode

import (
	"reflect"
	"testing"
)

func Test_Problem217(t *testing.T) {
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
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := containsDuplicate(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
