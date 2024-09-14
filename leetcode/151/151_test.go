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
				s: "the sky is blue",
			},
			want: "blue is sky the",
		},
		{
			name: "example 2",
			args: args{
				s: "  hello world  ",
			},
			want: "world hello",
		},
		{
			name: "a good   example",
			args: args{
				s: "a good   example",
			},
			want: "example good a",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := reverseWords(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
