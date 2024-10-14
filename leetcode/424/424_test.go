package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := characterReplacement(tt.args.s, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
