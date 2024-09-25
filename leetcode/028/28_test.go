package leetcode

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()

	type args struct {
		haystack string
		needle   string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				haystack: "sadbutsad", needle: "sad",
			},
			want: 0,
		},
		{
			name: "example 2",
			args: args{
				haystack: "leetcode", needle: "leeto",
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := strStr(tt.args.haystack, tt.args.needle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
