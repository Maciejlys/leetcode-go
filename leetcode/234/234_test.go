package leetcode

import (
	"reflect"
	"testing"

	"github.com/Maciejlys/leetcode-go/structures"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		root *ListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				root: structures.Ints2List([]int{1, 2, 2, 1}),
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				root: structures.Ints2List([]int{1, 2}),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isPalindrome(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
