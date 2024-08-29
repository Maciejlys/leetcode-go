package leetcode

import (
	"reflect"
	"testing"

	"github.com/Maciejlys/leetcode-go/structures"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "example 1",
			args: args{
				root: structures.Ints2TreeNode([]int{4, 2, 7, 1, 3, 6, 9}),
			},
			want: structures.Ints2TreeNode([]int{4, 7, 2, 9, 6, 3, 1}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
