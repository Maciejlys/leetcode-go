package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		s    string
		k    int
		fill byte
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				s:    "abcdefghi",
				k:    3,
				fill: 'x',
			},
			want: []string{"abc", "def", "ghi"},
		},
		{
			name: "example 2",
			args: args{
				s:    "abcdefghij",
				k:    3,
				fill: 'x',
			},
			want: []string{"abc", "def", "ghi", "jxx"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := divideString(tt.args.s, tt.args.k, tt.args.fill); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
