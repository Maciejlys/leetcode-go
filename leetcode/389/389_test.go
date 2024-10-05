package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "example 1",
			args: args{
				s: "abcd", t: "abcde",
			},
			want: 'e',
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := findTheDifference(tt.args.s, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
