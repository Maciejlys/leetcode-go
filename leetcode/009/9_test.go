package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				n: 121,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				n: -121,
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				n: 10,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isPalindrome(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
