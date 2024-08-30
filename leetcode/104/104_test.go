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
		want int
	}{
		{
			name: "example 1",
			args: args{
				root: structures.Ints2TreeNode([]int{3, 9, 20, 5}),
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := maxDepth(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
