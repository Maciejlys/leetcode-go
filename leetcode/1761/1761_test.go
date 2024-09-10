package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		word1 string
		word2 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				word1: "abc",
				word2: "pqr",
			},
			want: "apbqcr",
		},
		{
			name: "example 2",
			args: args{
				word1: "ab",
				word2: "pqrs",
			},
			want: "apbqrs",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := mergeAlternately(tt.args.word1, tt.args.word2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
