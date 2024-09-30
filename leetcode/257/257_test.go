package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				root: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}},
			},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "example 2",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: []string{"1"},
		},
		{
			name: "example 3",
			args: args{
				root: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 5}}},
			},
			want: []string{"1->2->6", "1->2->5", "1->3"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
