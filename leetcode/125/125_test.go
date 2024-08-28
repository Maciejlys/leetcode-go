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
		want bool
	}{
		{
			name: "example 1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				s: " ",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isPalindrome(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
