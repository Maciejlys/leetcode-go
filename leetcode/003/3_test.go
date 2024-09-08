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
		want int
	}{
		{
			name: "example 1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				s: "au",
			},
			want: 2,
		},
		{
			name: "example 4",
			args: args{
				s: "aab",
			},
			want: 2,
		},
		{
			name: "example 5",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
		{
			name: "example 6",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := lengthOfLongestSubstring(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
