package leetcode

import (
	"reflect"
	"testing"

	"github.com/Maciejlys/leetcode-go/structures"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		head *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				head: structures.Ints2List([]int{1, 2, 3, 4, 5}),
			},
			want: structures.Ints2List([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
