package leetcode

import (
	"reflect"
	"testing"

	"github.com/Maciejlys/leetcode-go/structures"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		list1 *ListNode
		list2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				list1: structures.Ints2List([]int{1, 2, 4}),
				list2: structures.Ints2List([]int{1, 3, 4}),
			},
			want: structures.Ints2List([]int{1, 1, 2, 3, 4, 4}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
