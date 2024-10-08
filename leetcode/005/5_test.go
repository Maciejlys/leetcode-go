package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "example 2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "example 3",
			args: args{
				s: "a",
			},
			want: "a",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := longestPalindrome(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
