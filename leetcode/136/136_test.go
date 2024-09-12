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
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := singleNumber(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

var nums = []int{4, 1, 2, 1, 2}

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumber(nums)
	}
}

func BenchmarkSingleNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumber2(nums)
	}
}
