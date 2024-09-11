package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		node1 *TreeNode
		node2 *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				node1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
				node2: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				node1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
				node2: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := leafSimilar2(tt.args.node1, tt.args.node2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
