package leetcode

import (
	"reflect"
	"testing"
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
				head: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "example 2",
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			name: "example 3",
			args: args{
				head: nil,
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
